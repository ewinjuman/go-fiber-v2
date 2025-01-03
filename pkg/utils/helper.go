package utils

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"syscall"
	"time"
)

func DoWithRetry(ctx context.Context, fn func() error) error {
	var lastErr error
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := fn(); err != nil {
				lastErr = err
				time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
				continue
			}
			return nil
		}
	}
	return fmt.Errorf("after 3 attempts: %w", lastErr)
}

func DoWithCustomRetry(ctx context.Context, fn func() error) error {
	backoff := []time.Duration{
		100 * time.Millisecond,
		200 * time.Millisecond,
		500 * time.Millisecond,
	}

	var lastErr error
	for i := 0; i < len(backoff); i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := fn(); err != nil {
				lastErr = err
				// Log error untuk debugging
				log.Printf("Attempt %d failed: %v", i+1, err)

				if i < len(backoff)-1 {
					// Tambah jitter ke backoff
					jitter := time.Duration(rand.Int63n(int64(backoff[i] / 2)))
					time.Sleep(backoff[i] + jitter)
					continue
				}
			}
			return nil
		}
	}

	return fmt.Errorf("after %d attempts: %w", len(backoff), lastErr)
}

type RetryableError struct {
	err error
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("retryable: %v", e.err)
}

func (e *RetryableError) Unwrap() error {
	return e.err
}

// Helper untuk menentukan apakah error perlu di-retry
func ShouldRetry(err error) bool {
	if err == nil {
		return false
	}

	// Check specific error types
	var netErr net.Error
	if errors.As(err, &netErr) {
		return netErr.Temporary() || netErr.Timeout()
	}

	// Check retryable error wrapper
	var retryErr *RetryableError
	if errors.As(err, &retryErr) {
		return true
	}

	// Common network/IO errors yang perlu di-retry
	if errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.ECONNABORTED) ||
		errors.Is(err, syscall.EPIPE) {
		return true
	}

	// HTTP status codes yang biasa di-retry
	if strings.Contains(err.Error(), "429 Too Many Requests") ||
		strings.Contains(err.Error(), "503 Service Unavailable") ||
		strings.Contains(err.Error(), "502 Bad Gateway") {
		return true
	}

	// Connection errors
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "broken pipe") ||
		strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "no such host") ||
		strings.Contains(errStr, "timeout")
}

package example

import (
	"context"
	"encoding/json"
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	Rest "github.com/ewinjuman/go-lib/http"
	Session "github.com/ewinjuman/go-lib/session"
	"github.com/pkg/errors"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/repository"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"
)

type (
	UserHttpService interface {
		TokenSessionValidation(request ValidateSessionRequest) (response ValidateSessionResponse, err error)
	}

	userHttp struct {
		session         *Session.Session
		ottoUsersRest   Rest.RestClient
		ottoUsersConfig configs.Ottouser
	}
)

// New Create new http request to Users first
func NewUserHttp(session *Session.Session) UserHttpService {
	return &userHttp{
		session:         session,
		ottoUsersRest:   Rest.New(configs.Config.Ottouser.Option),
		ottoUsersConfig: configs.Config.Ottouser,
	}
}

// TokenSessionValidation Execute Request
func (o *userHttp) TokenSessionValidation(request ValidateSessionRequest) (response ValidateSessionResponse, err error) {

	//How to create header:
	//exampleHeaders := http.Header{} //common way
	//or
	//exampleHeaders := o.ottoUsersRest.DefaultHeader("userORid", "password") //create header include basic auth

	//How to create query param:
	//exampleQueryParam := map[string]string{ //using map
	//	"customerId":         customerId,
	//	"insuranceProductId": strconv.Itoa(insuranceProductId),
	//}

	result, httpStatus, err := o.ottoUsersRest.Execute(o.session, o.ottoUsersConfig.Host, o.ottoUsersConfig.Path.TokenValidation, http.MethodPost, nil, request, nil, nil)
	//do something if err is not nil
	if err != nil {
		if Error.IsTimeout(err) {
			//.... do something if needed
		}
		return
	}
	//do something if http code is not 200
	if httpStatus != 200 {
		err = Error.New(httpStatus, repository.FailedStatus, "Error validation")
		return
	}

	//Bind Response Body
	json.Unmarshal(result, &response)

	//Do Something with response body if needed
	if response.Code != 200 {
		err = Error.New(response.Code, response.Status, response.Message)
		//if response.Code == 400 {
		//	err = Error.New(fiber.StatusNotFound, response.Status, "Not Found")
		//}
		return
	}
	return
}

func (o *userHttp) TokenSessionValidationWithRetry(request ValidateSessionRequest) (response *ValidateSessionResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//var result ValidateSessionResponse
	makeRequest := func() error {
		result, httpStatus, err := o.ottoUsersRest.Execute(o.session, o.ottoUsersConfig.Host, o.ottoUsersConfig.Path.TokenValidation, http.MethodPost, nil, request, nil, nil)
		//do something if err is not nil
		if err != nil {
			if Error.IsTimeout(err) {
				//.... do something if needed
			}
			return err
		}
		//do something if http code is not 200
		if httpStatus != 200 {
			err = Error.New(httpStatus, repository.FailedStatus, "Error validation")
			return err
		}

		//Bind Response Body
		json.Unmarshal(result, &response)

		//Do Something with response body if needed
		if response.Code != 200 {
			err = Error.New(response.Code, response.Status, response.Message)
			return err
		}
		return nil
	}
	err = o.doWithRetry(ctx, makeRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *userHttp) doWithRetry(ctx context.Context, fn func() error) error {
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

func (s *userHttp) doWithCustomRetry(ctx context.Context, fn func() error) error {
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
func shouldRetry(err error) bool {
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

// Circuit breaker implementation
type CircuitBreaker struct {
	failures    int32
	lastFailure time.Time
	threshold   int32
	timeout     time.Duration
}

//func (cb *CircuitBreaker) Execute(fn func() error) error {
//	if !cb.canTry() {
//		return fmt.Errorf("circuit breaker open")
//	}
//
//	err := fn()
//	if err != nil {
//		cb.recordFailure()
//		return err
//	}
//
//	cb.reset()
//	return nil
//}

// Rate limiter implementation
type RateLimiter struct {
	tokens chan struct{}
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-rl.tokens:
		return nil
	}
}

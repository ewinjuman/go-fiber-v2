package http

import "time"

type Options struct {
	Timeout   time.Duration `json:"timeout"`
	DebugMode bool          `json:"debugMode"`
	SkipTLS   bool          `json:"skipTLS"`
}

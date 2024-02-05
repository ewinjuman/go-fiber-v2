package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"os"
	"strings"
)

const (
	SuccessCode   = 200
	ContinueCode  = 100
	UndefinedCode = 500
)

// const for Status
const (
	SuccessStatus    = "SUCCESS"
	PendingStatus    = "PENDING"
	FailedStatus     = "FAILED"
	UndefinedStatus  = "FAILED"
	ContinueStatus   = "CONTINUE"
	UndefinedMessage = "Internal Server Error"
)

func New(errorCode int, status, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Status:    status,
		Message:   message,
	}
}

type ApplicationError struct {
	ErrorCode int
	Status    string
	Message   string
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func IsTimeout(err error) bool {
	if os.IsTimeout(err) {
		return true
	}

	st, ok := status.FromError(err)
	if !ok {
		return false
	}

	return st.Code() == codes.DeadlineExceeded
}

func GetCode(err error) int {
	if err == nil {
		return 200
	}

	if he, ok := err.(*ApplicationError); ok {
		return he.ErrorCode
	}

	return 500
}

func ParseError(err error) *ApplicationError {
	if err == nil {
		return nil
	}

	// Check grpc error
	if he, ok := status.FromError(err); ok {
		code := codeApplication(he.Code())
		if code == SuccessCode {
			return nil
		}
		return &ApplicationError{
			ErrorCode: code,
			Status:    FailedStatus,
			Message:   he.Message(),
		}
	}

	// Check application error
	if he, ok := err.(*ApplicationError); ok {
		return he
	}

	// Default error
	m := err.Error()
	sErr := strings.Split(err.Error(), "=")
	if len(sErr) > 0 {
		m = strings.TrimSpace(sErr[len(sErr)-1])
	}
	return &ApplicationError{
		ErrorCode: http.StatusInternalServerError,
		Status:    FailedStatus,
		Message:   m,
	}
}

// NewError creates a new Error instance with an optional message
func NewError(code int, status string, message ...string) error {
	err := &ApplicationError{
		ErrorCode: code,
		Status:    status,
		Message:   StatusMessage(code),
	}
	if len(message) > 0 {
		err.Message = message[0]
	}
	return err
}

func StatusMessage(status int) string {
	if status >= 0 && status < len(statusMessage) {
		if m := statusMessage[status]; m != "" {
			return m
		}
	}
	return UndefinedMessage
}

// NOTE: Keep this in sync with the status code list
var statusMessage = []string{
	100: "Continue",            // StatusContinue
	101: "Switching Protocols", // StatusSwitchingProtocols
	102: "Processing",          // StatusProcessing
	103: "Early Hints",         // StatusEarlyHints

	200: "OK",                            // StatusOK
	201: "Created",                       // StatusCreated
	202: "Accepted",                      // StatusAccepted
	203: "Non-Authoritative Information", // StatusNonAuthoritativeInformation
	204: "No Content",                    // StatusNoContent
	205: "Reset Content",                 // StatusResetContent
	206: "Partial Content",               // StatusPartialContent
	207: "Multi-Status",                  // StatusMultiStatus
	208: "Already Reported",              // StatusAlreadyReported
	226: "IM Used",                       // StatusIMUsed

	300: "Multiple Choices",   // StatusMultipleChoices
	301: "Moved Permanently",  // StatusMovedPermanently
	302: "Found",              // StatusFound
	303: "See Other",          // StatusSeeOther
	304: "Not Modified",       // StatusNotModified
	305: "Use Proxy",          // StatusUseProxy
	306: "Switch Proxy",       // StatusSwitchProxy
	307: "Temporary Redirect", // StatusTemporaryRedirect
	308: "Permanent Redirect", // StatusPermanentRedirect

	400: "Bad Request",                     // StatusBadRequest
	401: "Unauthorized",                    // StatusUnauthorized
	402: "Payment Required",                // StatusPaymentRequired
	403: "Forbidden",                       // StatusForbidden
	404: "Not Found",                       // StatusNotFound
	405: "Method Not Allowed",              // StatusMethodNotAllowed
	406: "Not Acceptable",                  // StatusNotAcceptable
	407: "Proxy Authentication Required",   // StatusProxyAuthRequired
	408: "Request Timeout",                 // StatusRequestTimeout
	409: "Conflict",                        // StatusConflict
	410: "Gone",                            // StatusGone
	411: "Length Required",                 // StatusLengthRequired
	412: "Precondition Failed",             // StatusPreconditionFailed
	413: "Request Entity Too Large",        // StatusRequestEntityTooLarge
	414: "Request URI Too Long",            // StatusRequestURITooLong
	415: "Unsupported Media Type",          // StatusUnsupportedMediaType
	416: "Requested Range Not Satisfiable", // StatusRequestedRangeNotSatisfiable
	417: "Expectation Failed",              // StatusExpectationFailed
	418: "I'm a teapot",                    // StatusTeapot
	421: "Misdirected Request",             // StatusMisdirectedRequest
	422: "Unprocessable Entity",            // StatusUnprocessableEntity
	423: "Locked",                          // StatusLocked
	424: "Failed Dependency",               // StatusFailedDependency
	425: "Too Early",                       // StatusTooEarly
	426: "Upgrade Required",                // StatusUpgradeRequired
	428: "Precondition Required",           // StatusPreconditionRequired
	429: "Too Many Requests",               // StatusTooManyRequests
	431: "Request Header Fields Too Large", // StatusRequestHeaderFieldsTooLarge
	451: "Unavailable For Legal Reasons",   // StatusUnavailableForLegalReasons

	500: "Internal Server Error",           // StatusInternalServerError
	501: "Not Implemented",                 // StatusNotImplemented
	502: "Bad Gateway",                     // StatusBadGateway
	503: "Service Unavailable",             // StatusServiceUnavailable
	504: "Gateway Timeout",                 // StatusGatewayTimeout
	505: "HTTP Version Not Supported",      // StatusHTTPVersionNotSupported
	506: "Variant Also Negotiates",         // StatusVariantAlsoNegotiates
	507: "Insufficient Storage",            // StatusInsufficientStorage
	508: "Loop Detected",                   // StatusLoopDetected
	510: "Not Extended",                    // StatusNotExtended
	511: "Network Authentication Required", // StatusNetworkAuthenticationRequired
}

var ErrDeadlineExceeded = DeadlineExceededError()

type deadlineExceededError struct {
	err     string
	timeout bool
}

func (e *deadlineExceededError) Error() string   { return e.err }
func (e *deadlineExceededError) Timeout() bool   { return e.timeout }
func (e *deadlineExceededError) Temporary() bool { return true }

func DeadlineExceededError(message ...string) error {
	defaultMessage := "context deadline exceeded (Client.Timeout exceeded while awaiting headers)"
	if len(message) > 0 {
		defaultMessage = message[0]
	}
	return &deadlineExceededError{
		err:     defaultMessage,
		timeout: true,
	}
}

var rpcCodeToApplicationCode = map[codes.Code]int{
	codes.OK:                 200,
	codes.Canceled:           406,
	codes.Unknown:            500,
	codes.InvalidArgument:    400,
	codes.DeadlineExceeded:   451,
	codes.NotFound:           404,
	codes.AlreadyExists:      409,
	codes.PermissionDenied:   403,
	codes.ResourceExhausted:  500,
	codes.FailedPrecondition: 500,
	codes.Aborted:            500,
	codes.OutOfRange:         413,
	codes.Unimplemented:      501,
	codes.Internal:           500,
	codes.Unavailable:        502,
	codes.DataLoss:           500,
	codes.Unauthenticated:    401,
}

func codeApplication(status codes.Code) int {
	if int(status) >= len(rpcCodeToApplicationCode) {
		return int(status)
	}
	m := rpcCodeToApplicationCode[status]
	if m == 0 {
		m = int(status)
	}
	return m
}

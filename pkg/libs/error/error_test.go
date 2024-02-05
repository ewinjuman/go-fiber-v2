package error

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Custom Error", args{err: NewError(451, "FAILED", "set error pending")}, 451},
		{"Error", args{err: errors.New("set error")}, 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCode(tt.args.err); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name       string
		args       args
		wantResult *ApplicationError
	}{
		{"Custom Error", args{err: NewError(451, "FAILED", "set error pending")}, &ApplicationError{
			ErrorCode: 451,
			Status:    FailedStatus,
			Message:   "set error pending",
		}},
		{"Error", args{err: errors.New("set error")}, &ApplicationError{
			ErrorCode: http.StatusInternalServerError,
			Status:    FailedStatus,
			Message:   "set error",
		}},
		{"Error nil", args{err: nil}, nil},
		{"Custom Error no message", args{err: NewError(451, "FAILED")}, &ApplicationError{
			ErrorCode: 451,
			Status:    FailedStatus,
			Message:   "Unavailable For Legal Reasons",
		}},
		{"Custom Error no message and no code", args{err: NewError(600, "FAILED")}, &ApplicationError{
			ErrorCode: 600,
			Status:    FailedStatus,
			Message:   UndefinedMessage,
		}},
		{"Custom Error v1", args{err: New(403, "FAILED", "error v1")}, &ApplicationError{
			ErrorCode: 403,
			Status:    FailedStatus,
			Message:   "error v1",
		}},
		{"rpc Error", args{err: status.Error(451, "error rpc gan")}, &ApplicationError{
			ErrorCode: 451,
			Status:    FailedStatus,
			Message:   "error rpc gan",
		}},
		{"rpc Error", args{err: status.Error(11, "error rpc gan")}, &ApplicationError{
			ErrorCode: 413,
			Status:    FailedStatus,
			Message:   "error rpc gan",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ParseError(tt.args.err); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseError() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		code    int
		status  string
		message []string
	}
	tests := []struct {
		name string
		args args
		want *ApplicationError
	}{
		{"Create new custom error", args{
			code:    400,
			status:  FailedStatus,
			message: []string{"Bad Request"},
		}, &ApplicationError{
			ErrorCode: 400,
			Status:    FailedStatus,
			Message:   "Bad Request",
		}},
		{"Create new custom error with no massage", args{
			code:    400,
			status:  FailedStatus,
			message: nil,
		}, &ApplicationError{
			ErrorCode: 400,
			Status:    FailedStatus,
			Message:   "Bad Request",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.code, tt.args.status, tt.args.message...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusMessage(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"available code", args{status: 200}, "OK"},
		{"unavailable code", args{status: 600}, UndefinedMessage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusMessage(tt.args.status); got != tt.want {
				t.Errorf("StatusMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		errorCode int
		status    string
		message   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Create new error",
			args{
				errorCode: 400,
				status:    "FAILED",
				message:   "Gagal",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.errorCode, tt.args.status, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

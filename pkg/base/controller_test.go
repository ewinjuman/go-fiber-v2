package base

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

type objectdata struct {
	Email string `json:"email"`
}

func TestBuildResponse(t *testing.T) {
	type args struct {
		data interface{}
		err  error
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			"success",
			args{
				data: &objectdata{Email: "email@g.com"},
				err:  nil,
			},
			&Response{
				Code:    200,
				Status:  "SUCCESS",
				Message: "",
				Data:    &objectdata{Email: "email@g.com"},
			},
		},
		{
			"error",
			args{
				data: nil,
				err:  errors.New("gagal"),
			},
			&Response{
				Code:    500,
				Status:  "FAILED",
				Message: "gagal",
				Data:    nil,
			},
		},
		{
			"error with data",
			args{
				data: &objectdata{Email: "email@g.com"},
				err:  errors.New("gagal"),
			},
			&Response{
				Code:    500,
				Status:  "FAILED",
				Message: "gagal",
				Data:    &objectdata{Email: "email@g.com"},
			},
		},
		{
			"error apps",
			args{
				data: nil,
				err:  Error.New(400, "FAIL", "gagal"),
			},
			&Response{
				Code:    400,
				Status:  "FAIL",
				Message: "gagal",
				Data:    nil,
			},
		},
		{
			"error apps with data",
			args{
				data: &objectdata{Email: "email@g.com"},
				err:  Error.New(401, "FAIL", "gagal"),
			},
			&Response{
				Code:    401,
				Status:  "FAIL",
				Message: "gagal",
				Data:    &objectdata{Email: "email@g.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildResponse(tt.args.data, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

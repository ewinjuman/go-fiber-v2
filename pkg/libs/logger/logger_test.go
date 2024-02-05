package logger

import (
	"reflect"
	"testing"
)

type object struct {
	Pin      string     `json:"pin"`
	Password string     `json:"password"`
	Token    string     `json:"token"`
	Name     string     `json:"name"`
	JsonData objectdata `json:"jsonData"`
}

type objectdata struct {
	Email string `json:"email"`
}

func TestLogger_MaskingJson(t *testing.T) {
	type fields struct {
		Options Options
	}
	f := fields{Options: Options{MaskingLogJsonPath: "pin|password|token|jsonData"}}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
	}{
		{"string masking", f, args{data: object{
			Pin:      "123456",
			Password: "passini",
			Token:    "oslkd3",
			Name:     "pulan",
			JsonData: objectdata{Email: "ewin"},
		}}, map[string]any{"jsonData": "***Mask JSON***", "name": "pulan", "password": "******", "pin": "******", "token": "******"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Options: tt.fields.Options,
			}
			if gotResult := l.MaskingJson(tt.args.data); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MaskingJson() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestLogger_MaskingJsonWithPath(t *testing.T) {
	type fields struct {
		Options Options
	}
	f := fields{}
	type args struct {
		data     interface{}
		jsonPath string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
	}{
		{"string masking", f, args{data: object{
			Pin:      "123456",
			Password: "passini",
			Token:    "oslkd3",
			Name:     "pulan",
			JsonData: objectdata{Email: "ewin"},
		}, jsonPath: "pin|password|token|jsonData"}, map[string]any{"jsonData": "***Mask JSON***", "name": "pulan", "password": "******", "pin": "******", "token": "******"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Options: tt.fields.Options,
			}
			if gotResult := l.MaskingJsonWithPath(tt.args.data, tt.args.jsonPath); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MaskingJsonWithPath() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

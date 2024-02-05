package example

import (
	"go-fiber-v2/pkg/repository"
	"testing"
)

func TestGetCode(t *testing.T) {
	type args struct {
		rc string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Success code", args{"00"}, repository.SuccessCode},
		{"Bad request code", args{"01"}, repository.BadRequestCode},
		{"Pending code", args{"06"}, repository.PendingCode},
		{"Undefined code", args{"66"}, repository.UndefinedCode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCode(tt.args.rc); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package repository

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"testing"
)

func TestSetError(t *testing.T) {
	type args struct {
		code    int
		message []string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			"Success Test",
			args{
				code:    200,
				message: []string{"success"},
			},
			false,
			"",
		},
		{
			"Error Test",
			args{
				code:    405,
				message: []string{"transaction failed"},
			},
			true,
			"transaction failed",
		},
		{
			"Error Test found",
			args{
				code: PendingCode,
			},
			true,
			PendingErr.Error(),
		},
		{
			"Error Undefined Test",
			args{
				code: 980,
			},
			true,
			UndefinedErr.Error()},
		{
			"Error Undefined with message Test",
			args{
				code:    980,
				message: []string{"error message"},
			},
			true,
			"error message"},
		{
			"Error Test found with message",
			args{
				code:    PendingCode,
				message: []string{"error message"},
			},
			true,
			"error message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetError(tt.args.code, tt.args.message...); (err != nil) != tt.wantErr || (err != nil && err.Error() != tt.wantErrMessage) {
				t.Errorf("SetError() error = %v, wantErr %v,  wantErrMessage %v", err, tt.wantErr, tt.wantErrMessage)
			}
		})
	}
}

func TestHandleMysqlError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Error Duplicate",
			args{err: &mysql.MySQLError{Number: 1062, Message: "Duplicate Entry"}},
			true,
		},
		{
			"Error Undefined",
			args{err: &mysql.MySQLError{Number: 11062, Message: "Error Undefined"}},
			true,
		},
		{
			"Error Not Found",
			args{err: errors.New("record not found")},
			true,
		},
		{
			"Error nil",
			args{err: nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HandleMysqlError(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("ConvMysqlErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

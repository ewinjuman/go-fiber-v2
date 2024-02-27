package repository

import "testing"

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

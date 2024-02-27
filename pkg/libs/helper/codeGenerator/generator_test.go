package codeGenerator

import "testing"

func TestRandomString(t *testing.T) {
	type args struct {
		size    int
		charset []rune
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{
			"AllCharset size 3",
			args{
				size:    3,
				charset: AllCharset,
			},
			3,
		},
		{
			"AllCharset size 6",
			args{
				size:    6,
				charset: AllCharset,
			},
			6,
		},
		{
			"Size 0",
			args{
				size:    0,
				charset: AllCharset,
			},
			0,
		},
		{
			"Charset Empty",
			args{
				size:    3,
				charset: nil,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(tt.args.size, tt.args.charset); len(got) != tt.wantLength {
				t.Errorf("RandomString() length= %v, want length %v", len(got), tt.wantLength)
			}
		})
	}
}

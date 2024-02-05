package configs

import "testing"

func Test_getConfigFilePath(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"local",
			args{env: "local"},
			"./resource/conf/config.local.json",
		},
		{
			"dev",
			args{env: "dev"},
			"./resource/conf/config.dev.json",
		},
		{
			"prod",
			args{env: "prod"},
			"./resource/conf/config.prod.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigFilePath(tt.args.env); got != tt.want {
				t.Errorf("getConfigFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

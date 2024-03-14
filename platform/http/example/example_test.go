package example

//
//import (
//	"go-fiber-v2/pkg/configs"
//	Rest "go-fiber-v2/pkg/libs/http"
//	Session "go-fiber-v2/pkg/libs/session"
//	"go-fiber-v2/platform/http/example"
//	"reflect"
//	"testing"
//)
//
//func TestNddew(t *testing.T) {
//	tests := []struct {
//		name       string
//		config     *ottoUsersHttp
//		wantConfig interface{}
//	}{
//		{
//			"New Http OC",
//			&ottoUsersHttp{
//				ottoUsersRest:   Rest.New(configs.Config.Ottouser.Option),
//				ottoUsersConfig: configs.Config.Ottouser,
//			},
//			configs.Config.Ottouser,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := New(); !reflect.DeepEqual(got.ottoUsersConfig, tt.wantConfig) {
//				t.Errorf("New() = %v, want %v", got.ottoUsersConfig, tt.wantConfig)
//			}
//		})
//	}
//}
//
//func TestNewUserHttp(t *testing.T) {
//	type args struct {
//		session *Session.Session
//	}
//	tests := []struct {
//		name string
//		args args
//		want example.UserHttpService
//	}{
//		{
//			"",
//			args{session: },
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := example.NewUserHttp(tt.args.session); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewUserHttp() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

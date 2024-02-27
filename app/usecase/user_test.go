package usecase

//
//import (
//	"github.com/stretchr/testify/mock"
//	"go-fiber-v2/app/models"
//	"go-fiber-v2/app/queries"
//	"go-fiber-v2/pkg/configs"
//	Error "go-fiber-v2/pkg/libs/error"
//	Logger "go-fiber-v2/pkg/libs/logger"
//	"go-fiber-v2/pkg/libs/session"
//	"go-fiber-v2/platform/http/example"
//	"reflect"
//	"testing"
//)
//
//func Test_userUsecase_CreateUser(t *testing.T) {
//	log := Logger.New(configs.Config.Logger)
//	mockUserQueriesSuccess := new(queries.MockUserQueries)
//	respQueryInsertUser := &models.User{
//		ID:           45,
//		Email:        "ew@email.com",
//		PasswordHash: "$2a$04$6/BNlaKd2HU4QoesEVzbRumnLfLwNxTcikjEkiXXUWjqfb29FDPQq",
//		UserStatus:   1,
//		UserRole:     "admin",
//	}
//	mockUserQueriesSuccess.On("InsertOneItem", mock.AnythingOfType("*models.User")).Return(respQueryInsertUser, nil)
//
//	mockUserQueriesError := new(queries.MockUserQueries)
//	mockUserQueriesError.On("InsertOneItem", mock.AnythingOfType("*models.User")).Return(&models.User{}, Error.NewError(400, "FAILED", "Duplicate Entry"))
//
//	mockUserHttpSuccess := new(example.MockUserHttp)
//	resHttpSuccess := example.ValidateSessionResponse{
//		Code:    200,
//		Status:  "Success",
//		Message: "Success",
//	}
//	resHttpSuccess.Data.OldID = 1
//	resHttpSuccess.Data.MobilePhoneNumber = "08123456"
//
//	mockUserHttpSuccess.On("TokenSessionValidation", mock.AnythingOfType("ValidateSessionRequest")).Return(resHttpSuccess, nil)
//
//	mockUserHttpError := new(example.MockUserHttp)
//	mockUserHttpError.On("TokenSessionValidation", mock.AnythingOfType("ValidateSessionRequest")).Return(example.ValidateSessionResponse{}, Error.NewError(400, "FAILED", "Duplicate Entry"))
//
//	type fields struct {
//		userQuery queries.UserQueriesService
//		userHttp  example.UserHttpService
//	}
//	type args struct {
//		session *session.Session
//		up      *models.SignUpRequest
//	}
//	tests := []struct {
//		name         string
//		fields       fields
//		args         args
//		wantResponse interface{}
//		wantErr      bool
//	}{
//		{
//			"Success test",
//			fields{
//				userQuery: mockUserQueriesSuccess,
//				userHttp:  mockUserHttpSuccess,
//			},
//			args{
//				session: session.New(log),
//				up: &models.SignUpRequest{
//					Email:    "ew@mail.com",
//					Password: "12345",
//					UserRole: "admin",
//				},
//			},
//			models.SignUpResponse{
//				ID:                respQueryInsertUser.ID,
//				Email:             respQueryInsertUser.Email,
//				Status:            respQueryInsertUser.UserStatus,
//				UserRole:          respQueryInsertUser.UserRole,
//				OldID:             resHttpSuccess.Data.OldID,
//				MobilePhoneNumber: resHttpSuccess.Data.MobilePhoneNumber,
//			},
//			false,
//		},
//		{
//			"Error test",
//			fields{
//				userQuery: mockUserQueriesError,
//				userHttp:  mockUserHttpSuccess,
//			},
//			args{
//				session: session.New(log),
//				up: &models.SignUpRequest{
//					Email:    "ew@mail.com",
//					Password: "12345",
//					UserRole: "admin",
//				},
//			},
//			nil,
//			true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := &userUsecase{
//				session:   tt.args.session,
//				userQuery: tt.fields.userQuery,
//				userHttp:  tt.fields.userHttp,
//			}
//			gotResponse, err := h.CreateUser(tt.args.up)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
//				t.Errorf("CreateUser() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
//			}
//		})
//	}
//}

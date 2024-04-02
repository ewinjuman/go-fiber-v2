package usecase

import (
	"github.com/stretchr/testify/mock"
	Error "gitlab.pede.id/otto-library/golang/share-pkg/error"
	Logger "gitlab.pede.id/otto-library/golang/share-pkg/logger"
	"gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/app/domain/entities"
	queries2 "go-fiber-v2/app/domain/queries"
	"go-fiber-v2/app/models"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/platform/grpc/user"
	"go-fiber-v2/platform/http/example"
	"reflect"
	"testing"
)

func Test_userUsecase_CreateUser(t *testing.T) {
	log := Logger.New(configs.Config.Logger)
	mockUserQueriesSuccess := new(queries2.MockUserQueries)
	respQueryInsertUser := &entities.User{
		ID:           45,
		Email:        "ew@email.com",
		PasswordHash: "$2a$04$6/BNlaKd2HU4QoesEVzbRumnLfLwNxTcikjEkiXXUWjqfb29FDPQq",
		UserStatus:   1,
		UserRole:     "admin",
	}
	mockUserQueriesSuccess.On("InsertOneItem", mock.AnythingOfType("*models.User")).Return(respQueryInsertUser, nil)

	mockUserQueriesError := new(queries2.MockUserQueries)
	mockUserQueriesError.On("InsertOneItem", mock.AnythingOfType("*models.User")).Return(&entities.User{}, Error.NewError(400, "FAILED", "Duplicate Entry"))

	mockUserHttpSuccess := new(example.MockUserHttp)
	resHttpSuccess := example.ValidateSessionResponse{
		Code:    200,
		Status:  "Success",
		Message: "Success",
	}
	resHttpSuccess.Data.OldID = 1
	resHttpSuccess.Data.MobilePhoneNumber = "08123456"

	mockUserHttpSuccess.On("TokenSessionValidation", mock.AnythingOfType("ValidateSessionRequest")).Return(resHttpSuccess, nil)

	mockUserHttpError := new(example.MockUserHttp)
	mockUserHttpError.On("TokenSessionValidation", mock.AnythingOfType("ValidateSessionRequest")).Return(example.ValidateSessionResponse{}, Error.NewError(400, "FAILED", "Duplicate Entry"))

	mockUserGrpcSuccess := new(user.MockUserGrpc)
	mockUserGrpcSuccess.On("TokenValidation", mock.AnythingOfType("string")).Return("081393746665", nil)

	mockUserGrpcError := new(user.MockUserGrpc)
	mockUserGrpcError.On("TokenValidation", mock.AnythingOfType("string")).Return("", Error.NewError(401, "FAILED", "Session anda telah habis"))

	type fields struct {
		session   *session.Session
		userQuery queries2.UserQueriesService
		userHttp  example.UserHttpService
		userGrpc  user.UserGrpcService
	}
	type args struct {
		up *models.SignUpRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse interface{}
		wantErr      bool
	}{
		{
			"Success test",
			fields{
				session:   session.New(log),
				userQuery: mockUserQueriesSuccess,
				userHttp:  mockUserHttpSuccess,
				userGrpc:  mockUserGrpcSuccess,
			},
			args{
				up: &models.SignUpRequest{
					Email:    "ew@mail.com",
					Password: "12345",
					UserRole: "admin",
				},
			},
			models.SignUpResponse{
				ID:                respQueryInsertUser.ID,
				Email:             respQueryInsertUser.Email,
				Status:            respQueryInsertUser.UserStatus,
				UserRole:          respQueryInsertUser.UserRole,
				OldID:             resHttpSuccess.Data.OldID,
				MobilePhoneNumber: resHttpSuccess.Data.MobilePhoneNumber,
			},
			false,
		},
		{
			"Error test",
			fields{
				session:   session.New(log),
				userQuery: mockUserQueriesError,
				userHttp:  mockUserHttpSuccess,
				userGrpc:  mockUserGrpcSuccess,
			},
			args{
				up: &models.SignUpRequest{
					Email:    "ew@mail.com",
					Password: "12345",
					UserRole: "admin",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &userUsecase{
				session:   tt.fields.session,
				userQuery: tt.fields.userQuery,
				userHttp:  tt.fields.userHttp,
				userGrpc:  tt.fields.userGrpc,
			}
			gotResponse, err := h.CreateUser(tt.args.up)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("CreateUser() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

package user

import (
	"context"
	"go-fiber-v2/app/models"
	"go-fiber-v2/app/usecase"
	Error "go-fiber-v2/pkg/libs/error"
	"go-fiber-v2/pkg/libs/helper/convert"
	Session "go-fiber-v2/pkg/libs/session"
	"go-fiber-v2/pkg/repository"
	"go-fiber-v2/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) SignUp(ctx context.Context, in *RequestSignUp) (*ResponseSignUp, error) {
	session := ctx.Value(Session.AppSession).(*Session.Session)
	request := &models.SignUpRequest{
		Email:    in.Email,
		Password: in.Password,
		UserRole: in.UserRole,
	}

	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		// Return, if some fields are not valid.
		return nil, status.Error(codes.Code(repository.BadRequestCode), err.Error())
	}

	user := usecase.NewUserUsecase(session)
	d, err := user.CreateUser(request)
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	result := &models.SignUpResponse{}
	convert.ObjectToObject(d, result)
	response := &ResponseSignUp{
		Id:                int32(result.ID),
		Email:             result.Email,
		Status:            int32(result.Status),
		UserRole:          result.UserRole,
		OldId:             int32(result.OldID),
		MobilePhoneNumber: result.MobilePhoneNumber,
	}

	return response, nil
}

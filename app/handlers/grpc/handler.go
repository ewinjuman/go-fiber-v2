package grpc

import (
	"context"
	Error "gitlab.pede.id/otto-library/golang/share-pkg/error"
	"gitlab.pede.id/otto-library/golang/share-pkg/helper/convert"
	Session "gitlab.pede.id/otto-library/golang/share-pkg/session"
	pb2 "go-fiber-v2/app/handlers/grpc/pb"
	"go-fiber-v2/app/models"
	"go-fiber-v2/app/usecase/user"
	"go-fiber-v2/pkg/repository"
	"go-fiber-v2/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb2.UnimplementedUserServer
}

func (s *server) SignUp(ctx context.Context, in *pb2.RequestSignUp) (*pb2.ResponseSignUp, error) {
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

	user := user.NewUserUsecase(session)
	d, err := user.CreateUser(request)
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	result := &models.SignUpResponse{}
	convert.ObjectToObject(d, result)
	response := &pb2.ResponseSignUp{
		Id:                int32(result.ID),
		Email:             result.Email,
		Status:            int32(result.Status),
		UserRole:          result.UserRole,
		OldId:             int32(result.OldID),
		MobilePhoneNumber: result.MobilePhoneNumber,
	}

	return response, nil
}

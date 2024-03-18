package usecase

import (
	"gitlab.pede.id/otto-library/golang/share-pkg/password"
	"gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/app/models"
	"go-fiber-v2/app/queries"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/platform/grpc/user"
	"go-fiber-v2/platform/http/example"
)

type (
	UserUsecaseService interface {
		CreateUser(up *models.SignUpRequest) (interface{}, error)
	}

	userUsecase struct {
		session   *session.Session
		userGrpc  user.UserGrpcService
		userQuery queries.UserQueriesService
		userHttp  example.UserHttpService
	}
)

func NewUserUsecase(session *session.Session) (item UserUsecaseService) {
	//userGrpc, err := user.NewServerContext()
	return &userUsecase{
		session:   session,
		userGrpc:  user.NewUserGrpc(session),
		userQuery: queries.NewUserQueries(session),
		userHttp:  example.NewUserHttp(session),
	}
}

func (h *userUsecase) CreateUser(up *models.SignUpRequest) (response interface{}, err error) {
	user := &models.User{}

	user.Email = up.Email
	user.PasswordHash = password.GeneratePassword(up.Password)
	user.UserStatus = 1
	user.UserRole = up.UserRole + configs.Config.Apps.Name

	//Grpc Example
	//resultUserGrpc, err := h.userGrpc.TokenValidation("tokenUser")
	//if err != nil {
	//	return
	//}
	//println(resultUserGrpc)

	//SQL query example
	resultQuery, err := h.userQuery.InsertOneItem(user)
	if err != nil {
		return
	}

	//http example
	reqHttp := example.ValidateSessionRequest{Token: "tokenUser"}
	respHttp, err := h.userHttp.TokenSessionValidation(reqHttp)
	if err != nil {
		return
	}

	response = models.SignUpResponse{
		ID:                resultQuery.ID,
		Email:             resultQuery.Email,
		Status:            resultQuery.UserStatus,
		UserRole:          resultQuery.UserRole,
		OldID:             respHttp.Data.OldID,
		MobilePhoneNumber: respHttp.Data.MobilePhoneNumber,
	}
	return
}

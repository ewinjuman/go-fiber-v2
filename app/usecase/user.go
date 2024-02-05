package usecase

import (
	"go-fiber-v2/app/models"
	"go-fiber-v2/app/queries"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/libs/session"
	"go-fiber-v2/pkg/utils"
	"go-fiber-v2/platform/http/example"
)

type (
	UserUsecaseService interface {
		CreateUser(up *models.SignUpRequest) (interface{}, error)
	}

	userUsecase struct {
		session   *session.Session
		userQuery queries.UserQueriesService
		userHttp  example.UserHttpService
	}
)

func NewUserUsecase(session *session.Session) (item UserUsecaseService) {
	return &userUsecase{
		session:   session,
		userQuery: queries.NewUserQueries(session),
		userHttp:  example.NewUserHttp(session),
	}
}

func (h *userUsecase) CreateUser(up *models.SignUpRequest) (response interface{}, err error) {
	user := &models.User{}

	user.Email = up.Email
	user.PasswordHash = utils.GeneratePassword(up.Password)
	user.UserStatus = 1
	user.UserRole = up.UserRole + configs.Config.Apps.Name

	resultQuery, err := h.userQuery.InsertOneItem(user)
	if err != nil {
		return
	}

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

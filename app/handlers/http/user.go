package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	Error "gitlab.pede.id/otto-library/golang/share-pkg/error"
	"go-fiber-v2/app/models"
	"go-fiber-v2/app/usecase/user"
	"go-fiber-v2/pkg/base"
	"go-fiber-v2/pkg/repository"
	"go-fiber-v2/pkg/utils"
)

func UserSignUp(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	
	// Create a new user auth struct.
	signUp := &models.SignUpRequest{}

	// Checking received data from JSON body.
	if err := ctx.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}
	user := user.NewUserUsecase(ctx.Session)
	result, err := user.CreateUser(signUp)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func UserSignOut(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	return ctx.Response("User sign Out", nil)
}

func UserSignUpV2(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	return ctx.Response("User sign Up v2", nil)
}

func Generate(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	panic("panic")
	return ctx.Response(uuid.New().ID(), nil)
}

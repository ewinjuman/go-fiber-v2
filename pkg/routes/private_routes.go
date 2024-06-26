package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "go-fiber-v2/app/handlers/http"
	"go-fiber-v2/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Delete("/user/sign/out", middleware.JWTProtected(), httpHandler.UserSignOut) // de-authorization user
}

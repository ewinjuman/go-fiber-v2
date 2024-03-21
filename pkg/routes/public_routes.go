package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	httpHandler "go-fiber-v2/app/handlers/http"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")
	v1 := route.Group("/v1")
	v1.Get("/monitor", monitor.New(monitor.Config{Title: "Go-Fiber v2 Metrics Page"}))
	v1.Get("/health", httpHandler.HealthCheck) // register a new user
	// Routes for POST method:
	v1.Post("/user/sign/up", httpHandler.UserSignUp) // register a new user
	v1.Get("/generate", httpHandler.Generate)

	v2 := route.Group("/v2")
	v2.Post("/user/sign/up", httpHandler.UserSignUpV2)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go-fiber-v2/app/httpHandler"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	route.Get("/monitor", monitor.New(monitor.Config{Title: "Go-Fiber v2 Metrics Page"}))
	route.Get("/health", httpHandler.HealthCheck) // register a new user
	// Routes for POST method:
	route.Post("/user/sign/up", httpHandler.UserSignUp) // register a new user
}

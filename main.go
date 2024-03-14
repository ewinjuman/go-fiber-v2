package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	userGrpc "go-fiber-v2/app/interfaces/grpc"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/middleware"
	"go-fiber-v2/pkg/routes"
	"go-fiber-v2/pkg/utils"
)

// @title Skeleton Service API
// @version 2.0
// @description Skeleton service using golang and fiber framework.
// @Still continuing to develop
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private_libs.sh routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	myFigure := figure.NewColorFigure("GAMBIT - FIBER v2", "", "green", true)
	myFigure.Print()

	go func() {
		userGrpc.StartGrpcServer()
	}()
	// Start server (with or without graceful shutdown).
	if configs.Config.Apps.Mode == "local" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

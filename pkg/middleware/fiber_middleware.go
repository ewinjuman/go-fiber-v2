package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	Logger "gitlab.pede.id/otto-library/golang/share-pkg/logger"
	Session "gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/pkg/configs"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add request response logger.
		RequestResponseLog,

		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodHead,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodPatch,
			}, ","),
		}),

		//Add panic recovery
		recover.New(recover.Config{EnableStackTrace: true, StackTraceHandler: stackTraceHandler}),

		// Add idempotency
		idempotency.New(idempotency.Config{
			Lifetime: 30 * time.Minute,
			// ...
		}),

		//Add Rewrite for backward compatibility or just creating cleaner and more descriptive links
		//rewrite.New(rewrite.Config{
		//	Rules: map[string]string{
		//		"/api/v1/user/sign/up": "/api/v2/user/sign/up",
		//	},
		//}),

		// Add simple logger.
		//logger.New(),

	)
}

// RequestResponseLog logging for request and response API
func RequestResponseLog(c *fiber.Ctx) error {
	// Create new log
	log := Logger.New(configs.Config.Logger)

	uri := c.OriginalURL()

	var request interface{}
	json.Unmarshal(c.Body(), &request)

	session := Session.New(log).
		SetInstitutionID(configs.Config.Apps.DefaultAppsId).
		SetAppName(configs.Config.Apps.Name).
		SetURL(uri).
		SetMethod(c.Method()).
		SetRequest(request).
		SetHeader(c.GetReqHeaders()).
		SetActionTo("FE to BE")

	if uri != "/api/v1/monitor" {
		session.LogRequest("Log Request")
	}

	c.Context().SetUserValue(Session.AppSession, session)

	// Go to next handler:
	c.Next()

	// Log response
	var response interface{}
	if err := json.Unmarshal(c.Context().Response.Body(), &response); err != nil {
		return err
	}

	if uri != "/api/v1/monitor" {
		session.LogResponse(response, "Log Response")
	}

	return nil
}

func stackTraceHandler(c *fiber.Ctx, err interface{}) {
	s := Session.GetSession(c)
	s.Error(err)
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", err, debug.Stack()))
	// Return status 500 and failed authentication error.
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": "Internal Server Error",
		"status":  "ERROR",
		"data":    nil,
	})
}

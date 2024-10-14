package base

import (
	Error "github.com/ewinjuman/go-lib/error"
	Session "github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	"go-fiber-v2/pkg/repository"
)

type Base struct {
	*fiber.Ctx
	Session *Session.Session
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewContext method to parse *fiber.Ctx and get Session.
func NewContext(c *fiber.Ctx) Base {
	return Base{
		Ctx:     c,
		Session: Session.GetSession(c),
	}
}

// Response Method to create response api.
func (b *Base) Response(data interface{}, err error) error {
	return b.Ctx.JSON(BuildResponse(data, err))
}

func BuildResponse(data interface{}, err error) *Response {
	res := &Response{
		Data: data,
	}

	if err != nil {
		if he, ok := err.(*Error.ApplicationError); ok {
			res.Code = he.ErrorCode
			res.Status = he.Status
		} else {
			res.Code = repository.UndefinedCode
			res.Status = repository.UndefinedStatus
		}
		res.Message = err.Error()
	} else {
		res.Code = repository.SuccessCode
		res.Status = repository.SuccessStatus
	}

	return res
}

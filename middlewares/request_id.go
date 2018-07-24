package middlewares

import (
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

const (
	RequestIDContextKey = "request_id_context_key"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		uuid, err := uuid.NewV4()
		if err != nil {
			c.Logger().Error(err)
		}
		c.Set(RequestIDContextKey, uuid)
		return next(c)
	})
}

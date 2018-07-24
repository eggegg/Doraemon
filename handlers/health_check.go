package handlers

import (
	"net/http"

	"github.com/eggegg/Doraemon/middlewares"

	"github.com/eggegg/Doraemon/renderings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func HealthCheck(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
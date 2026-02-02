package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(app *echo.Echo, msgs string) {
	app.GET("/health-check", func(c echo.Context) error {
		type HealthCheck struct {
			Message string `json:"message"`
		}
		return c.JSON(
			http.StatusOK,
			&HealthCheck{
				Message: msgs,
			},
		)
	})
}

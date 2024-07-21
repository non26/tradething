package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlerwareConposing(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
}

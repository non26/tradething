package main

import (
	"fmt"
	"tradething/cmd/app"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	app_echo := echo.New()
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, config)
	app.RouteSemiBotComposing(app_echo, config)

	app_echo.Start(fmt.Sprintf(":%v", config.Port))
}

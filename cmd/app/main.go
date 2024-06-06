package main

import (
	"fmt"
	"net/http"
	route "tradething/cmd/app/route/app"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		type HealthCheck struct {
			Message string `json:"message"`
		}
		return c.JSON(
			http.StatusOK,
			&HealthCheck{
				Message: "success",
			},
		)
	})

	route.BkRouting(
		app,
		config.Secrets.BitkubApikey,
		config.Secrets.BitkubSecretKey,
		config.ServiceName.KubSpot,
		&config.KubSpotUrl,
	)
	route.BnRouting(
		app,
		config.ServiceName.BinanceFuture,
		&config.Secrets,
		&config.BinanceFutureUrl)
	route.KcRouting(
		app,
		config.ServiceName.KucoinSpot,
		config.ServiceName.KucoinFuture,
		config.Secrets.KucoinApiKey,
		config.Secrets.KucoinApiKeyVersion,
		config.Secrets.KucoinSecretKey,
		config.Secrets.KucoinPassphase,
		&config.KCFutureUrl,
		&config.KCSpotUrl,
	)
	route.OkxRouting(
		app,
		config.Secrets.OkxApiKey,
		config.Secrets.OkxSecretKey,
		config.Secrets.OkxPassPhase,
		config.ServiceName.OKXFuture,
		&config.OkxFutureUrl,
		&config.Secrets,
		config.Env,
	)

	app.Start(fmt.Sprintf(":%v", config.Port))
}

package route

import (
	bnservice "tradething/app/bn/app/bnservice/future"
	"tradething/config"
	timeinterval "tradething/semibot/timeinterval/bn"

	"github.com/labstack/echo/v4"
)

func TimeIntervalRoute(
	app *echo.Echo,
	config *config.AppConfig,
) {
	service_name := "bn-timeinterval-semibot"

	bnservice := bnservice.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
	)
	semibot := timeinterval.NewBnTimeIntervalService(bnservice)
	handler := timeinterval.NewBnTradeTimeIntervalHandler(
		config,
		service_name,
		semibot,
	)

	bnTimeIntervalGroup := app.Group("/" + service_name)
	bnTimeIntervalGroup.POST("/bn-timeinterval-semibot", handler.BnHandler)
}

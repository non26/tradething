package route

import (
	bnservice "tradething/app/bn/app/bnservice/future"
	handler "tradething/app/bn/app/handler/future"
	service "tradething/app/bn/app/service/future"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func BnRouting(
	app *echo.Echo,
	serviceName string,
	secret *config.Secrets,
	bnFutureConfig *config.BinanceFutureUrl,
) {
	binanceGroup := app.Group("/" + serviceName)
	binanceServie := bnservice.NewBinanceFutureExternalService(
		bnFutureConfig,
		secret,
		serviceName,
	)
	service := service.NewBinanceFutureService(
		serviceName,
		binanceServie,
	)

	placeOrderHandler := handler.NewPlaceSinglerOrderHandler(
		service,
	)
	binanceGroup.POST("/place-order", placeOrderHandler.Handler)

	setLeverageHandler := handler.NewsetNewLeveragehandler(
		service,
	)
	binanceGroup.POST("/set-leverage", setLeverageHandler.Handler)

	queryOrderHandler := handler.NewqueryOrderHandler(
		service,
	)
	binanceGroup.POST("/query-order", queryOrderHandler.Handler)
}

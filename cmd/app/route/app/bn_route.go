package route

import (
	bnservice "tradetoolv2/app/bn/app/bnservice/future"
	handler "tradetoolv2/app/bn/app/handler/future"
	service "tradetoolv2/app/bn/app/service/future"
	"tradetoolv2/config"

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
}

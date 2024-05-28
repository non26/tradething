package route

import (
	"tradetoolv2/app/okx/app/handler"
	"tradetoolv2/app/okx/app/okxservice"
	"tradetoolv2/app/okx/app/service"
	"tradetoolv2/config"

	"github.com/labstack/echo/v4"
)

func OkxRouting(
	app *echo.Echo,
	okxApi string,
	okxSecret string,
	okxPassPhase string,
	serviceName string,
	okxConfig *config.OkxFutureUrl, // spot+future
	secrets *config.Secrets,
	env string,
) {
	okxGroup := app.Group("/" + serviceName)
	okxService := okxservice.NewOKXExternalService(
		okxConfig,
		secrets,
		env,
	)
	service := service.NewOkxService(
		okxService,
	)

	placeSingleOrderHandler := handler.NewPlaceSinglePositionHandler(
		service,
	)
	okxGroup.POST("/place-single-order", placeSingleOrderHandler.Handler)

	placeMultiOrderHandler := handler.NewPlaceMultiplePositionHandler(
		service,
	)
	okxGroup.POST("/place-multi-order", placeMultiOrderHandler.Handler)

	setLeverageOrderHandler := handler.NewSetLeverage(
		service,
	)
	okxGroup.POST("/set-leverage", setLeverageOrderHandler.Handler)
}

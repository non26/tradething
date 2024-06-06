package route

import (
	futurehandler "tradething/app/kc/app/handler/future/order"
	spothandler "tradething/app/kc/app/handler/spot/order"
	kcfutureservice "tradething/app/kc/app/kcservice/future/order"
	kcspotservice "tradething/app/kc/app/kcservice/spot/order"
	futureservice "tradething/app/kc/app/service/future/order"
	spotservice "tradething/app/kc/app/service/spot/order"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func KcRouting(
	app *echo.Echo,
	spotServiceName string,
	futureServiceName string,
	kcApi string,
	kcApiVersion string,
	kcSecret string,
	kcPassPhase string,
	kcFutureConfig *config.KCFutureUrl,
	kcSpotConfig *config.KCSpotUrl,
) {

	kcSpotRouting(
		app,
		spotServiceName,
		kcApi,
		kcApiVersion,
		kcSecret,
		kcPassPhase,
		kcSpotConfig,
	)
	kcFutureRouting(
		app,
		futureServiceName,
		kcApi,
		kcApiVersion,
		kcSecret,
		kcPassPhase,
		kcFutureConfig,
	)

}

func kcSpotRouting(
	app *echo.Echo,
	serviceName string,
	kcApi string,
	kcApiVersion string,
	kcSecret string,
	kcPassPhase string,
	kcSpotConfig *config.KCSpotUrl,
) {
	kucoincSpotGroup := app.Group("/" + serviceName)
	kucoinSpotService := kcspotservice.NewSpotKcService(
		kcApi,
		kcApiVersion,
		kcSecret,
		kcPassPhase,
		serviceName,
		kcSpotConfig,
	)
	service := spotservice.NewSpotOrderService(
		kucoinSpotService,
	)

	placeOrderHandler := spothandler.NewPlaceSpotORderHandler(
		service,
	)
	kucoincSpotGroup.POST("/place-order", placeOrderHandler.Handler)

}

func kcFutureRouting(
	app *echo.Echo,
	serviceName string,
	kcApi string,
	kcApiVersion string,
	kcSecret string,
	kcPassPhase string,
	kcFutureConfig *config.KCFutureUrl,
) {
	kucoincFutureGroup := app.Group("/" + serviceName)
	kucoinFutureService := kcfutureservice.NewFutureKcService(
		kcApi,
		kcApiVersion,
		kcSecret,
		kcPassPhase,
		serviceName,
		kcFutureConfig,
	)
	service := futureservice.NewFutureOrderService(
		kucoinFutureService,
	)

	placeOrderHandler := futurehandler.NewPlaceSpotORderHandler(
		service,
	)
	kucoincFutureGroup.POST("place-order", placeOrderHandler.Handler)

}

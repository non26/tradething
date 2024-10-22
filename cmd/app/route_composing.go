package app

import (
	bnservice "tradething/app/bn/bn_future/bnservice"
	bothandler "tradething/app/bn/bn_future/bot_handler"
	botservice "tradething/app/bn/bn_future/bot_service"
	handler "tradething/app/bn/bn_future/handler"
	svcrepository "tradething/app/bn/bn_future/repository"
	service "tradething/app/bn/bn_future/service"
	"tradething/app/bn/bncommon"
	lambdaroute "tradething/cmd/app/route/lambda"

	"tradething/config"

	"github.com/labstack/echo/v4"
)

func RouteRestApiConposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType bncommon.IOrderType,
	positionSide bncommon.IPositionSide,
	side bncommon.ISide,
	svcRepository svcrepository.IRepository,
	httpttransport bncommon.IBinanceServiceHttpTransport,
	httpclient bncommon.IBinanceSerivceHttpClient,
) {
	binanceGroup := app.Group("/" + config.ServiceName.BinanceFuture)
	binanceServie := bnservice.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
		httpttransport,
		httpclient,
	)
	service := service.NewBinanceFutureService(
		config.ServiceName.BinanceFuture,
		binanceServie,
		svcRepository,
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

func RouteSemiBotComposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType bncommon.IOrderType,
	positionSide bncommon.IPositionSide,
	side bncommon.ISide,
	svcRepository svcrepository.IRepository,
	httpttransport bncommon.IBinanceServiceHttpTransport,
	httpclient bncommon.IBinanceSerivceHttpClient,
) {
	service_name := "bn-future-semibot"

	bnservice := bnservice.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
		httpttransport,
		httpclient,
	)
	botservice := botservice.NewBotService(bnservice, svcRepository, orderType, positionSide, side)

	bothandler := bothandler.NewBotHandler(config, service_name, botservice)
	bnTimeIntervalGroup := app.Group("/" + service_name)
	bnTimeIntervalGroup.POST("/time-frame-interval", bothandler.BotTimeFrameIntervalHandler)
}

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	lambdaroute.UpdateAWSAppConfig(app, config)
}

package app

import (
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	bothandler "tradething/app/bn/bn_future/bot_handler"
	botservice "tradething/app/bn/bn_future/bot_service"
	handler "tradething/app/bn/bn_future/handler"
	service "tradething/app/bn/bn_future/service"
	lambdaroute "tradething/cmd/app/route/lambda"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconst "github.com/non26/tradepkg/pkg/bn/position_constant"

	"tradething/config"

	"github.com/labstack/echo/v4"
)

func RouteRestApiComposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType positionconst.IOrderType,
	positionSide positionconst.IPositionSide,
	side positionconst.ISide,
	svcRepository bndynamodb.IRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
	bnTradeService bntrade.IBinanceFutureExternalService,
	bnMarketService bnmarket.IBnMarketDataService,
) {
	binanceGroup := app.Group("/" + config.ServiceName.BinanceFuture)
	service := service.NewBinanceFutureService(
		config.ServiceName.BinanceFuture,
		bnTradeService,
		bnMarketService,
		svcRepository,
		orderType,
		positionSide,
		side,
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

	placeMultipleOrderHandler := handler.NewPlaceMultipleOrderHandler(
		service,
	)
	binanceGroup.POST("/place-multiple-order", placeMultipleOrderHandler.Handler)

	// close order by multiple client id
	closeByClientIdsHandler := handler.NewCloseByClientIdsHandler(
		service,
	)
	binanceGroup.POST("/close-by-client-ids", closeByClientIdsHandler.Handler)

	setPositionHandler := handler.NewSetPositionHandler(
		service,
	)
	binanceGroup.POST("/set-position", setPositionHandler.Handler)
}

func RouteBotRestApiComposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType positionconst.IOrderType,
	positionSide positionconst.IPositionSide,
	side positionconst.ISide,
	svcRepository bndynamodb.IRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
	bnTradeService bntrade.IBinanceFutureExternalService,
	bnMarketService bnmarket.IBnMarketDataService,
) {
	botGroup := app.Group("/" + config.ServiceName.BinanceFuture + "/bot")
	botService := botservice.NewBotService(
		bnTradeService,
		svcRepository,
		orderType,
		positionSide,
		side,
	)

	botTimeframeExeIntervalHandler := bothandler.NewBotTimeframeExeIntervalHandler(
		botService,
	)
	botGroup.POST("/timeframe-exe-interval", botTimeframeExeIntervalHandler.Handler)

	invalidateBotHandler := bothandler.NewInvalidateBotHandler(
		botService,
	)
	botGroup.POST("/invalidate", invalidateBotHandler.Handler)
}

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	lambdaroute.UpdateAWSAppConfig(app, config)
}

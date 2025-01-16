package app

import (
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	botservice "tradething/app/bn/bn_future/bot"
	bothandler "tradething/app/bn/bn_future/bot_handler"
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
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPositionTable bndynamodb.IBnFtAdvancedPositionRepository,
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
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
		bnFtAdvancedPositionTable,
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

	invalidatePositionHandler := handler.NewInvalidatePositionHandler(
		service,
	)
	binanceGroup.POST("/invalidate-position", invalidatePositionHandler.Handler)

	// set-advanced-position
}

func RouteBotRestApiComposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType positionconst.IOrderType,
	positionSide positionconst.IPositionSide,
	side positionconst.ISide,
	bnFtBotTable bndynamodb.IBnFtBotRepository,
	bnFtBotOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
	bnTradeService bntrade.IBinanceFutureExternalService,
	bnMarketService bnmarket.IBnMarketDataService,
) {
	botGroup := app.Group("/" + config.ServiceName.BinanceFuture + "/bot")
	botService := botservice.NewBotService(
		bnTradeService,
		bnFtBotTable,
		bnFtBotOnRunTable,
		bnFtHistoryTable,
		bnFtQouteUsdtTable,
		orderType,
		positionSide,
		side,
	)

	botTimeframeExeIntervalHandler := bothandler.NewBotTimeframeExeIntervalHandler(
		botService,
	)
	botGroup.POST("/timeframe-exe-interval", botTimeframeExeIntervalHandler.Handler)

	// invalidateBotHandler := bothandler.NewInvalidateBotHandler(
	// 	botService,
	// )
	// botGroup.POST("/invalidate", invalidateBotHandler.Handler)
}

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	lambdaroute.UpdateAWSAppConfig(app, config)
}

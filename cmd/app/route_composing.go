package app

import (
	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"
	handler "tradething/app/bn/bn_future/handler"
	service "tradething/app/bn/bn_future/service"
	lambdaroute "tradething/cmd/app/route/lambda"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"

	"tradething/config"

	"github.com/labstack/echo/v4"
)

func RouteRestApiComposing(
	app *echo.Echo,
	config *config.AppConfig,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtQouteUSDTRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
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
	)

	placeOrderHandler := handler.NewPlaceSinglerOrderHandler(
		service,
	)
	binanceGroup.POST("/position", placeOrderHandler.Handler)

	setLeverageHandler := handler.NewsetNewLeveragehandler(
		service,
	)
	binanceGroup.POST("/new-leverage", setLeverageHandler.Handler)

	queryOrderHandler := handler.NewqueryOrderHandler(
		service,
	)
	binanceGroup.POST("/query-order", queryOrderHandler.Handler)

	placeMultipleOrderHandler := handler.NewPlaceMultipleOrderHandler(
		service,
	)
	binanceGroup.POST("/positions", placeMultipleOrderHandler.Handler)

	closeByClientIdsHandler := handler.NewCloseByClientIdsHandler(
		service,
	)
	binanceGroup.POST("/close-by-client-ids", closeByClientIdsHandler.Handler)

	invalidatePositionHandler := handler.NewInvalidatePositionHandler(
		service,
	)
	binanceGroup.POST("/invalidate-position", invalidatePositionHandler.Handler)
}

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	lambdaroute.UpdateAWSAppConfig(app, config)
}

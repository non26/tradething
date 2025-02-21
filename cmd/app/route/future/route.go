package route

import (
	handlers "tradething/app/bn/handlers/future"
	adaptor "tradething/app/bn/infrastructure/adaptor/future"
	infratrade "tradething/app/bn/infrastructure/future"
	infraposition "tradething/app/bn/infrastructure/future/position"
	process "tradething/app/bn/process/future"

	"tradething/config"

	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

func RouteFuture(
	app *echo.Echo,
	config *config.AppConfig,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtQouteUsdtTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) {
	binanceGroup := app.Group("/" + config.ServiceName.BinanceFuture)
	ftAdaptor := adaptor.NewBinanceFutureAdaptorService(
		&config.BinanceFutureUrl,
		config.Secrets.BinanceApiKey,
		config.Secrets.BinanceSecretKey,
		config.ServiceName.BinanceFuture,
		httpttransport,
		httpclient,
	)

	longPosition := infraposition.NewLongPosition(
		ftAdaptor,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
	)

	shortPosition := infraposition.NewShortPosition(
		ftAdaptor,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
	)

	trade := infratrade.NewTrade(
		longPosition,
		shortPosition,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
	)

	ftProcess := process.NewFuture(
		trade,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
	)

	positionHandler := handlers.NewPositionHandler(ftProcess)
	binanceGroup.POST("/position", positionHandler.Handler)

	multiPositionHandler := handlers.NewMultiplePositionHandler(ftProcess)
	binanceGroup.POST("/positions", multiPositionHandler.Handler)

	closeByIdsHandler := handlers.NewCloseByIdHandler(ftProcess)
	binanceGroup.POST("/close-by-ids", closeByIdsHandler.Handler)

}

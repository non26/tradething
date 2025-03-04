package route

import (
	handlers "tradething/app/bn/handlers/future"
	adaptor "tradething/app/bn/infrastructure/adaptor/future"
	infra "tradething/app/bn/infrastructure/future"
	infraposition "tradething/app/bn/infrastructure/future/position"
	process "tradething/app/bn/process/future"

	"tradething/config"

	savepositionby "tradething/app/bn/infrastructure/future/save"

	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

func RouteFuture(
	app *echo.Echo,
	config *config.AppConfig,
	bnFtOpeningPositionTable bndynamodb.IBnFtOpeningPositionRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
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
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	shortPosition := infraposition.NewShortPosition(
		ftAdaptor,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	tradeBuilder := infra.NewTradePosition(
		longPosition,
		shortPosition,
	)

	trade := infra.NewTrade(
		tradeBuilder,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	lookUp := infra.NewTradeLookUp(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	saveBuyPosition := savepositionby.NewSaveBuyPosition(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	saveSellPosition := savepositionby.NewSaveSellPosition(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	savePositionBuilder := infra.NewSavePositionBuilder(
		saveBuyPosition,
		saveSellPosition,
	)

	savePosition := infra.NewSavePosition(savePositionBuilder)

	closePositionLookUp := infra.NewClosePositionLookUp(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	ftProcess := process.NewFuture(
		trade,
		lookUp,
		savePosition,
		closePositionLookUp,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	positionHandler := handlers.NewPositionHandler(ftProcess)
	binanceGroup.POST("/position", positionHandler.Handler)

	multiPositionHandler := handlers.NewMultiplePositionHandler(ftProcess)
	binanceGroup.POST("/positions", multiPositionHandler.Handler)

	closeByIdsHandler := handlers.NewCloseByIdHandler(ftProcess)
	binanceGroup.POST("/close-by-ids", closeByIdsHandler.Handler)

}

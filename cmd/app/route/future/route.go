package route

import (
	handlers "tradething/app/bn/handlers/future"
	adaptor "tradething/app/bn/infrastructure/adaptor/future"
	infraBuilder "tradething/app/bn/infrastructure/future/builder"
	infraLookup "tradething/app/bn/infrastructure/future/lookup"
	infraposition "tradething/app/bn/infrastructure/future/position"
	infraSave "tradething/app/bn/infrastructure/future/save"
	infraSavePosition "tradething/app/bn/infrastructure/future/save_position"
	infraTrade "tradething/app/bn/infrastructure/future/trade"
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
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtAdvancedPosition bndynamodb.IBnFtAdvancedPositionRepository,
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

	tradeBuilder := infraBuilder.NewTradePosition(
		longPosition,
		shortPosition,
	)

	trade := infraTrade.NewTrade(
		tradeBuilder,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	tradeLookUp := infraLookup.NewTradeLookUp(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	saveBuyPosition := infraSave.NewSaveBuyPosition(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	saveSellPosition := infraSave.NewSaveSellPosition(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
	)

	savePositionBuilder := infraBuilder.NewSavePositionBuilder(
		saveBuyPosition,
		saveSellPosition,
	)

	savePosition := infraSavePosition.NewSavePosition(savePositionBuilder)

	advancedPositionLookUp := infraLookup.NewAdvancedPositionLookUp(
		bnFtOpeningPositionTable,
		bnFtHistoryTable,
		bnFtAdvancedPosition,
	)

	cryptoLookUp := infraLookup.NewCryptoLookUp(
		bnFtCryptoTable,
	)

	process := process.NewFuture(
		trade,
		savePosition,
		tradeLookUp,
		advancedPositionLookUp,
		cryptoLookUp,
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
		bnFtAdvancedPosition,
	)

	positionHandler := handlers.NewPositionHandler(process)
	binanceGroup.POST("/position", positionHandler.Handler)

	multiPositionHandler := handlers.NewMultiplePositionHandler(process)
	binanceGroup.POST("/positions", multiPositionHandler.Handler)

	// manage-position
	closeByIdsHandler := handlers.NewCloseByIdHandler(process)
	binanceGroup.POST("/close-by-ids", closeByIdsHandler.Handler)

	// set-advanced-position
	// get-advanced-position

}

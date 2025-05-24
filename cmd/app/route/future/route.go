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
	apphandler "github.com/non26/tradepkg/pkg/bn/app_handler"
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
		bnFtAdvancedPosition,
	)

	saveSellPosition := infraSave.NewSaveSellPosition(
		bnFtOpeningPositionTable,
		bnFtCryptoTable,
		bnFtHistoryTable,
		bnFtAdvancedPosition,
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
	binanceGroup.POST("/position", apphandler.HandlerWrapper(positionHandler))

	multiPositionHandler := handlers.NewMultiplePositionHandler(process)
	binanceGroup.POST("/positions", apphandler.HandlerWrapper(multiPositionHandler))

	// close-position
	closeByIdsHandler := handlers.NewCloseByIdHandler(process)
	binanceGroup.POST("/close-by-ids", apphandler.HandlerWrapper(closeByIdsHandler))

	// set-advanced-position
	setAdvancedPositionHandler := handlers.NewSetAdvancedPositionHandler(process)
	binanceGroup.POST("/set-advanced-position", apphandler.HandlerWrapper(setAdvancedPositionHandler))

	// get-advanced-position
	getAdvancedPositionHandler := handlers.NewGetAdvancedPositionHandler(process)
	binanceGroup.POST("/get-advanced-position", apphandler.HandlerWrapper(getAdvancedPositionHandler))

	//manage-position
	mangePositionHandler := handlers.NewManagePositionHandler(process)
	binanceGroup.POST("/manage-position", apphandler.HandlerWrapper(mangePositionHandler))

}

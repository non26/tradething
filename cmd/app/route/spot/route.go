package route

import (
	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"

	handlers "tradething/app/bn/handlers/spot"
	adaptor "tradething/app/bn/infrastructure/adaptor/spot"
	infrastructure "tradething/app/bn/infrastructure/spot"
	infraorder "tradething/app/bn/infrastructure/spot/order"
	infrasave "tradething/app/bn/infrastructure/spot/save"
	process "tradething/app/bn/process/spot"
	"tradething/config"
)

func RouteSpot(
	app *echo.Echo,
	config *config.AppConfig,
	bnSpotOpeningPositionTable bndynamodb.IBnSpotOpeningPositionRepository,
	bnSpotQouteUsdtTable bndynamodb.IBnSpotCryptoRepository,
	bnSpotHistoryTable bndynamodb.IBnSpotHistoryRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) {
	spotGroup := app.Group("/" + config.ServiceName.BinanceSpot)
	spotAdaptor := adaptor.NewBinanceSpotAdaptorService(
		&config.BinanceSpotUrl,
		config.Secrets.BinanceSpotApiKey,
		config.Secrets.BinanceSpotSecretKey,
		config.ServiceName.BinanceSpot,
		httpttransport,
		httpclient,
	)

	infraSpotLookUp := infrastructure.NewTradeLookUp(
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	saveBuy := infrasave.NewSaveBuy(
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	saveSell := infrasave.NewSaveSell(
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	saveOrderBuilder := infrastructure.NewTradeSaveOrderBuilder(
		saveBuy,
		saveSell,
	)

	infraSaveOrder := infrastructure.NewTradeSaveOrder(
		saveOrderBuilder,
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	infraClosePositionLookUp := infrastructure.NewCloseOrderLookUp(
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	orderSpot := infraorder.NewOrderSpot(
		spotAdaptor,
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	tradeSpot := infrastructure.NewSpotTrade(
		orderSpot,
	)

	process := process.NewSpot(
		infraSpotLookUp,
		infraSaveOrder,
		infraClosePositionLookUp,
		tradeSpot,
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	spotHandler := handlers.NewOrderHandler(process)
	spotGroup.POST("/order", spotHandler.Handler)

	multipleOrderHandler := handlers.NewMultipleOrderHandler(process)
	spotGroup.POST("/orders", multipleOrderHandler.Handler)

	closeByIdHandler := handlers.NewCloseByIdHandler(process)
	spotGroup.POST("/close-by-ids", closeByIdHandler.Handler)

}

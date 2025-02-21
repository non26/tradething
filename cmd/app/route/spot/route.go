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

	orderSpot := infraorder.NewOrderSpot(
		spotAdaptor,
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
	)

	tradeSpot := infrastructure.NewSpotTrade(
		orderSpot,
	)

	spotProcess := process.NewSpot(
		bnSpotOpeningPositionTable,
		bnSpotQouteUsdtTable,
		bnSpotHistoryTable,
		tradeSpot,
	)

	spotHandler := handlers.NewOrderHandler(spotProcess)
	spotGroup.POST("/order", spotHandler.Handler)

	multipleOrderHandler := handlers.NewMultipleOrderHandler(spotProcess)
	spotGroup.POST("/orders", multipleOrderHandler.Handler)

	closeByIdHandler := handlers.NewCloseByIdHandler(spotProcess)
	spotGroup.POST("/close-by-ids", closeByIdHandler.Handler)

}

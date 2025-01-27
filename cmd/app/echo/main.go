package main

import (
	"fmt"
	"net/http"
	"tradething/cmd/app"

	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"

	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

func main() {
	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	bnFtOpeningPositionTable := bndynamodb.NewConnectionBnFtOpeningPositionRepository(dynamodbclient)
	bnFtQouteUsdtTable := bndynamodb.NewConnectionBnFtQouteUSDTRepository(dynamodbclient)
	bnFtHistoryTable := bndynamodb.NewConnectionBnFtHistoryRepository(dynamodbclient)
	bnFtBotTable := bndynamodb.NewConnectionBnFtBotRepository(dynamodbclient)
	bnFtBotOnRunTable := bndynamodb.NewConnectionBnFtBotOnRunRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	marketData := bnmarket.NewBnMarketDataService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)
	binanceServie := bntrade.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		config.Secrets.BinanceApiKey,
		config.Secrets.BinanceSecretKey,
		config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)

	bot_binanceServie := bntrade.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		config.Secrets.BinanceSubAccountApikey,
		config.Secrets.BinanceSubAccountSecretKey,
		config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)

	app_echo := echo.New()
	app.HealthCheck(app_echo)
	app.RouteRestApiComposing(
		app_echo,
		config,
		bnFtOpeningPositionTable,
		bnFtQouteUsdtTable,
		bnFtHistoryTable,
		httptransport,
		httpclient,
		binanceServie,
		marketData,
	)
	app.RouteBotRestApiComposing(
		app_echo,
		config,
		bnFtBotTable,
		bnFtBotOnRunTable,
		bnFtHistoryTable,
		bnFtQouteUsdtTable,
		httptransport,
		httpclient,
		bot_binanceServie,
		marketData,
	)
	app_echo.Start(fmt.Sprintf(":%v", config.Port))
}

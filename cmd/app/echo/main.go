package main

import (
	"net/http"
	"tradething/cmd/app"
	routefuture "tradething/cmd/app/route/future"
	routespot "tradething/cmd/app/route/spot"

	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	bndynamodbspot "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

func main() {
	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	// dynamodb config
	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	// dynamodb for future
	bnFtOpeningPositionTable := bndynamodb.NewConnectionBnFtOpeningPositionRepository(dynamodbclient)
	bnFtQouteUsdtTable := bndynamodb.NewConnectionBnFtCryptoRepository(dynamodbclient)
	bnFtHistoryTable := bndynamodb.NewConnectionBnFtHistoryRepository(dynamodbclient)
	bnFtAdvancedPosition := bndynamodb.NewConnectionBnFtAdvancedPositionRepository(dynamodbclient)
	// spot
	bnSpotOpeningPositionTable := bndynamodbspot.NewConnectionBnSpotOpeningPositionRepository(dynamodbclient)
	bnSpotQouteUsdtTable := bndynamodbspot.NewConnectionBnSpotCryptoRepository(dynamodbclient)
	bnSpotHistoryTable := bndynamodbspot.NewConnectionBnSpotHistoryRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()
	// echo
	app_echo := echo.New()
	app.HealthCheck(app_echo)
	// route
	routefuture.RouteFuture(app_echo, config, bnFtOpeningPositionTable, bnFtQouteUsdtTable, bnFtHistoryTable, bnFtAdvancedPosition, httptransport, httpclient)
	routespot.RouteSpot(app_echo, config, bnSpotOpeningPositionTable, bnSpotQouteUsdtTable, bnSpotHistoryTable, httptransport, httpclient)
	app_echo.Start(config.GetPortWithFormat())
}

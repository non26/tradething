package main

import (
	"context"
	"net/http"
	"tradething/cmd/app"
	"tradething/config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"

	bnmarket "tradething/app/bn/bn_future/bnservice/market_data"
	bntrade "tradething/app/bn/bn_future/bnservice/trade"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.AppConfig

func init() {

	var err error
	_config, err = app.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(_config.Dynamodb.Region, _config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(_config.Dynamodb.Ak, _config.Dynamodb.Sk)
	var dynamodbclient *dynamodb.Client
	if _config.IsLocal() {
		dynamodbclient = bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	} else {
		dynamodbclient = bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	}
	bnFtOpeningPositionTable := bndynamodb.NewConnectionBnFtOpeningPositionRepository(dynamodbclient)
	bnFtQouteUsdtTable := bndynamodb.NewConnectionBnFtQouteUSDTRepository(dynamodbclient)
	bnFtHistoryTable := bndynamodb.NewConnectionBnFtHistoryRepository(dynamodbclient)
	bnFtBotTable := bndynamodb.NewConnectionBnFtBotRepository(dynamodbclient)
	bnFtBotOnRunTable := bndynamodb.NewConnectionBnFtBotOnRunRepository(dynamodbclient)
	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	marketData := bnmarket.NewBnMarketDataService(
		&_config.BinanceFutureUrl,
		&_config.Secrets,
		_config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)
	binanceServie := bntrade.NewBinanceFutureExternalService(
		&_config.BinanceFutureUrl,
		_config.Secrets.BinanceApiKey,
		_config.Secrets.BinanceSecretKey,
		_config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)

	bot_binanceServie := bntrade.NewBinanceFutureExternalService(
		&_config.BinanceFutureUrl,
		_config.Secrets.BinanceSubAccountApikey,
		_config.Secrets.BinanceSubAccountSecretKey,
		_config.ServiceName.BinanceFuture,
		httptransport,
		httpclient,
	)
	app_echo := echo.New()
	app.MiddlerwareComposing(app_echo)
	app.HealthCheck(app_echo)
	app.RouteRestApiComposing(
		app_echo,
		_config,
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
		_config,
		bnFtBotTable,
		bnFtBotOnRunTable,
		bnFtHistoryTable,
		bnFtQouteUsdtTable,
		httptransport,
		httpclient,
		bot_binanceServie,
		marketData,
	)
	app.RouteLambda(app_echo, _config)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

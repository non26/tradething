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

	routefuture "tradething/cmd/app/route/future"
	routelambda "tradething/cmd/app/route/lambda"
	routespot "tradething/cmd/app/route/spot"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
	bndynamodbspot "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.AppConfig

func init() {

	var err error
	_config, err = app.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	// dynamodb config
	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(_config.Dynamodb.Region, _config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(_config.Dynamodb.Ak, _config.Dynamodb.Sk)
	var dynamodbclient *dynamodb.Client
	if _config.IsLocal() {
		dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	} else {
		dynamodbclient = bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	}
	// dynamodb for future
	bnFtOpeningPositionTable := bndynamodb.NewConnectionBnFtOpeningPositionRepository(dynamodbclient)
	bnFtQouteUsdtTable := bndynamodb.NewConnectionBnFtCryptoRepository(dynamodbclient)
	bnFtHistoryTable := bndynamodb.NewConnectionBnFtHistoryRepository(dynamodbclient)
	// dynamodb for spot
	bnSpotOpeningPositionTable := bndynamodbspot.NewConnectionBnSpotOpeningPositionRepository(dynamodbclient)
	bnSpotQouteUsdtTable := bndynamodbspot.NewConnectionBnSpotCryptoRepository(dynamodbclient)
	bnSpotHistoryTable := bndynamodbspot.NewConnectionBnSpotHistoryRepository(dynamodbclient)
	// http
	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	// echo
	app_echo := echo.New()
	app.MiddlerwareComposing(app_echo)
	app.HealthCheck(app_echo)
	// route
	routefuture.RouteFuture(app_echo, _config, bnFtOpeningPositionTable, bnFtQouteUsdtTable, bnFtHistoryTable, httptransport, httpclient)
	routespot.RouteSpot(app_echo, _config, bnSpotOpeningPositionTable, bnSpotQouteUsdtTable, bnSpotHistoryTable, httptransport, httpclient)
	routelambda.RouteLambda(app_echo, _config)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

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

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconst "github.com/non26/tradepkg/pkg/bn/position_constant"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.AppConfig

func init() {

	var err error
	_config, err = app.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	ordertype := positionconst.NewOrderType()
	side := positionconst.NewSide()
	positionSide := positionconst.NewPositionSide()
	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(_config.Dynamodb.Region, _config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(_config.Dynamodb.Ak, _config.Dynamodb.Sk)
	var dynamodbclient *dynamodb.Client
	if _config.IsLocal() {
		dynamodbclient = bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	} else {
		dynamodbclient = bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()
	}
	svcrepository := bndynamodb.NewDynamoDBRepository(dynamodbclient)
	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	app_echo := echo.New()
	app.MiddlerwareComposing(app_echo)
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, _config, ordertype, positionSide, side, svcrepository, httptransport, httpclient)
	// app.RouteSemiBotComposing(app_echo, _config, ordertype, positionSide, side, svcrepository, httptransport, httpclient)
	app.RouteLambda(app_echo, _config)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

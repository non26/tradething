package main

import (
	"context"
	"net/http"
	"tradething/app/bn/bncommon"
	"tradething/cmd/app"
	"tradething/config"

	svcrepository "tradething/app/bn/bn_future/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda
var _config *config.AppConfig

func init() {

	var err error
	_config, err = app.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	ordertype := bncommon.NewOrderType()
	side := bncommon.NewSide()
	positionSide := bncommon.NewPositionSide()
	dynamodbconfig := svcrepository.NewDynamodbConfig()
	dynamodbendpoint := svcrepository.NewEndPointResolver(_config.Dynamodb.Region, _config.Dynamodb.Endpoint)
	dynamodbcredential := svcrepository.NewCredential(_config.Dynamodb.Ak, _config.Dynamodb.Sk)
	dynamodbclient := svcrepository.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).New()
	svcrepository := svcrepository.NewDynamoDBRepository(dynamodbclient)
	httptransport := bncommon.NewBinanceTransport(&http.Transport{})
	httpclient := bncommon.NewBinanceSerivceHttpClient()

	app_echo := echo.New()
	app.MiddlerwareComposing(app_echo)
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, _config, ordertype, positionSide, side, svcrepository, httptransport, httpclient)
	app.RouteSemiBotComposing(app_echo, _config, ordertype, positionSide, side, svcrepository, httptransport, httpclient)
	app.RouteLambda(app_echo, _config)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

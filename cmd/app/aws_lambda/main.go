package main

import (
	"context"
	"tradething/cmd/app"
	"tradething/config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"

	route "tradething/cmd/app/route/future"

	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
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

	// echo
	app_echo := echo.New()
	app.MiddlerwareComposing(app_echo)
	app.HealthCheck(app_echo)
	// route
	route.RouteFuture(app_echo, _config, dynamodbclient)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

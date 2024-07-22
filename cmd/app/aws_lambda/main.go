package main

import (
	"context"
	"tradething/cmd/app"
	"tradething/config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func init() {

	// config, err := app.ReadLog("./opt")
	// if err != nil {
	// 	panic(fmt.Sprintf("This Lambda Function config not found %v", err.Error()))
	// }

	var config *config.AppConfig

	app_echo := echo.New()

	app.MiddlerwareConposing(app_echo)
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, config)
	app.RouteSemiBotRestApiConposing(app_echo, config)

	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

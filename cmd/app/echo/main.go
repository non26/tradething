package main

import (
	"tradething/cmd/app"
	route "tradething/cmd/app/route/future"

	"github.com/labstack/echo/v4"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	// bndynamodbspot "github.com/non26/tradepkg/pkg/bn/dynamodb_spot"
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

	// echo
	app_echo := echo.New()
	app.HealthCheck(app_echo)
	// route
	route.RouteFuture(app_echo, config, dynamodbclient)

	app_echo.Start(config.GetPortWithFormat())
}

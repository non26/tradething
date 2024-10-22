package main

import (
	"fmt"
	"tradething/cmd/app"

	svcrepository "tradething/app/bn/bn_future/repository"
	"tradething/app/bn/bncommon"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	ordertype := bncommon.NewOrderType()
	side := bncommon.NewSide()
	positionSide := bncommon.NewPositionSide()
	dynamodbconfig := svcrepository.NewDynamodbConfig()
	dynamodbendpoint := svcrepository.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := svcrepository.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := svcrepository.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).New()
	svcrepository := svcrepository.NewDynamoDBRepository(dynamodbclient)

	app_echo := echo.New()
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, config, ordertype, positionSide, side, svcrepository)
	app.RouteSemiBotComposing(app_echo, config, ordertype, positionSide, side, svcrepository)

	app_echo.Start(fmt.Sprintf(":%v", config.Port))
}

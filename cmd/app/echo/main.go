package main

import (
	"fmt"
	"net/http"
	"tradething/cmd/app"

	bnclient "github.com/non26/tradepkg/pkg/bn/binance_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/binance_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
	positionconst "github.com/non26/tradepkg/pkg/bn/position_constant"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	ordertype := positionconst.NewOrderType()
	side := positionconst.NewSide()
	positionSide := positionconst.NewPositionSide()
	dynamodbconfig := bndynamodb.NewDynamodbConfig()
	dynamodbendpoint := bndynamodb.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodb.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodb.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()
	svcrepository := bndynamodb.NewDynamoDBRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	app_echo := echo.New()
	app.HealthCheck(app_echo)
	app.RouteRestApiConposing(app_echo, config, ordertype, positionSide, side, svcrepository, httptransport, httpclient)

	app_echo.Start(fmt.Sprintf(":%v", config.Port))
}

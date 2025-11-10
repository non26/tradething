package route

import (
	marketdataroute "tradething/app/bn/future/market_data/route"
	positionroute "tradething/app/bn/future/position/route"
	positionhistoryroute "tradething/app/bn/future/position_history/route"
	subaccountroute "tradething/app/bn/future/sub_account/route"
	"tradething/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func RouteFuture(
	app_echo *echo.Echo,
	config *config.AppConfig,
	dynamodbclient *dynamodb.Client,
) {
	positionroute.Router(app_echo, dynamodbclient)
	marketdataroute.Router(app_echo, config)
	subaccountroute.Router(app_echo, dynamodbclient)
	positionhistoryroute.Router(app_echo, dynamodbclient)
}

package route

import (
	accmanagementroute "tradething/app/bn/future/BFF/account_management/router"
	accumroute "tradething/app/bn/future/accumulation/route"
	marketdataroute "tradething/app/bn/future/market_data/route"
	positionroute "tradething/app/bn/future/position/route"
	positionhistoryroute "tradething/app/bn/future/position_history/route"
	subaccountroute "tradething/app/bn/future/sub_account/route"
	"tradething/config"

	registeraccrepocoreservice "tradething/app/bn/future/sub_account/infrastructure/db/register_account"
	subaccountcoreservice "tradething/app/bn/future/sub_account/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func RouteFuture(
	app_echo *echo.Echo,
	config *config.AppConfig,
	dynamodbclient *dynamodb.Client,
) {
	regsiterAccountRepositoryCoreServce := registeraccrepocoreservice.NewRegisterAccountRepository(dynamodbclient)
	subaccountCoreService := subaccountcoreservice.NewSubAccountService(regsiterAccountRepositoryCoreServce)
	accmanagementroute.Router(app_echo, subaccountCoreService)

	if config.IsLocal() {
		positionroute.Router(app_echo, dynamodbclient)
		marketdataroute.Router(app_echo, config)
		subaccountroute.Router(app_echo, dynamodbclient)
		positionhistoryroute.Router(app_echo, dynamodbclient)
		accumroute.NewRoute(app_echo, dynamodbclient)
	}
}

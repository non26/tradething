package route

import (
	accumroute "tradething/app/bn/future/accumulation/route"
	marketdataroute "tradething/app/bn/future/market_data/route"
	positionroute "tradething/app/bn/future/position/route"
	positionhistoryroute "tradething/app/bn/future/position_history/route"
	subaccountroute "tradething/app/bn/future/sub_account/route"
	"tradething/config"

	accmanagementroute "tradething/app/bn/future/BFF/account_management/route"
	registeraccrepocoreservice "tradething/app/bn/future/sub_account/infrastructure/db/register_account"
	subaccountcoreservice "tradething/app/bn/future/sub_account/service"

	marketdatecoreserviceroute "tradething/app/bn/future/BFF/market_data/route"
	adaptor "tradething/app/bn/future/market_data/infrastructure/adaptor/market_data"
	marketdatacoreservice "tradething/app/bn/future/market_data/service"

	advpositioncoreserviceroute "tradething/app/bn/future/BFF/advanced_position/route"
	advpositionrepocoreservice "tradething/app/bn/future/position/infrastructure/db/adanced_position"
	advpositiocoreservice "tradething/app/bn/future/position/service/advanced_position"

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

	marketdataAdaptorCoreService := adaptor.NewMarketDataAdaptor(config.BinanceAdaptorFutureUsdt.BaseUrl, config.BinanceAdaptorFutureUsdt.KlinesCandleStick)
	marketdataCoreService := marketdatacoreservice.NewService(marketdataAdaptorCoreService)
	marketdatecoreserviceroute.Router(app_echo, marketdataCoreService)

	advPositionRepositoryCoreService := advpositionrepocoreservice.NewAdvancedPositionRepository(dynamodbclient)
	advPositionCoreService := advpositiocoreservice.NewAdvancedPositionService(advPositionRepositoryCoreService)
	advpositioncoreserviceroute.RegisterRoutes(app_echo, advPositionCoreService)

	if config.IsLocal() {
		positionroute.Router(app_echo, dynamodbclient)
		marketdataroute.Router(app_echo, config)
		subaccountroute.Router(app_echo, dynamodbclient)
		positionhistoryroute.Router(app_echo, dynamodbclient)
		accumroute.NewRoute(app_echo, dynamodbclient)
	}
}

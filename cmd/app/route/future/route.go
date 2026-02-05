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

	positionmanagementbffserviceroute "tradething/app/bn/future/BFF/position_mangement/route"
	currentpositionrepocoreservice "tradething/app/bn/future/position/infrastructure/db/opening_position"
	currentpositioncoreservice "tradething/app/bn/future/position/service/current_position"
	positionhistoryrepocoreservice "tradething/app/bn/future/position_history/infrastructure/db/history"
	positionhistorycoreservice "tradething/app/bn/future/position_history/service"

	tradebffserviceroute "tradething/app/bn/future/BFF/trade/route"
	accumulationrepocoreservice "tradething/app/bn/future/accumulation/infrastructure/db/accumulation"
	accumulationcoreservice "tradething/app/bn/future/accumulation/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func RouteFuture(
	app_echo *echo.Echo,
	config *config.AppConfig,
	dynamodbclient *dynamodb.Client,
) {
	groupBff := app_echo.Group("/bff")
	regsiterAccountRepositoryCoreServce := registeraccrepocoreservice.NewRegisterAccountRepository(dynamodbclient)
	subaccountCoreService := subaccountcoreservice.NewSubAccountService(regsiterAccountRepositoryCoreServce)
	accmanagementroute.Router(groupBff, subaccountCoreService)

	marketdataAdaptorCoreService := adaptor.NewMarketDataAdaptor(config.BinanceAdaptorFutureUsdt.BaseUrl, config.BinanceAdaptorFutureUsdt.KlinesCandleStick)
	marketdataCoreService := marketdatacoreservice.NewService(marketdataAdaptorCoreService)
	marketdatecoreserviceroute.Router(groupBff, marketdataCoreService)

	advPositionRepositoryCoreService := advpositionrepocoreservice.NewAdvancedPositionRepository(dynamodbclient)
	advPositionCoreService := advpositiocoreservice.NewAdvancedPositionService(advPositionRepositoryCoreService)
	advpositioncoreserviceroute.RegisterRoutes(groupBff, advPositionCoreService)

	positionHistoryRepositoryCoreService := positionhistoryrepocoreservice.NewBnFtHistoryRepository(dynamodbclient)
	positionHistoryCoreService := positionhistorycoreservice.NewService(positionHistoryRepositoryCoreService)
	currentPositionRepositoryCoreService := currentpositionrepocoreservice.NewOpeningPositionRepository(dynamodbclient)
	currentPositionCoreService := currentpositioncoreservice.NewCurrentPositionService(currentPositionRepositoryCoreService)
	positionmanagementbffserviceroute.Router(groupBff, currentPositionCoreService, positionHistoryCoreService)

	accumulationRepoCoreService := accumulationrepocoreservice.NewBnFtAccumulationRepository(dynamodbclient)
	accumulationCoreService := accumulationcoreservice.NewBnFtAccumulationService(accumulationRepoCoreService)
	tradebffserviceroute.Route(groupBff, config, accumulationCoreService, currentPositionCoreService, advPositionCoreService, positionHistoryCoreService, subaccountCoreService)

	groupCore := app_echo.Group("/core")
	if config.IsLocal() {
		positionroute.Router(groupCore, dynamodbclient)
		marketdataroute.Router(groupCore, config)
		subaccountroute.Router(groupCore, dynamodbclient)
		positionhistoryroute.Router(groupCore, dynamodbclient)
		accumroute.NewRoute(groupCore, dynamodbclient)
	}
}

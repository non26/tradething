package app

import (
	svcrepository "tradething/app/bn/bn_future/repository"
	"tradething/app/bn/bncommon"
	route "tradething/cmd/app/route/app"
	semibotroute "tradething/cmd/app/route/bot_route"
	lambdaroute "tradething/cmd/app/route/lambda"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func RouteRestApiConposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType bncommon.IOrderType,
	positionSide bncommon.IPositionSide,
	side bncommon.ISide,
	svcRepository svcrepository.IRepository) {
	route.BnRouting(
		app,
		config.ServiceName.BinanceFuture,
		&config.Secrets,
		&config.BinanceFutureUrl,
		orderType,
		positionSide,
		side,
		svcRepository)

}

func RouteSemiBotComposing(
	app *echo.Echo,
	config *config.AppConfig,
	orderType bncommon.IOrderType,
	positionSide bncommon.IPositionSide,
	side bncommon.ISide,
	svcRepository svcrepository.IRepository,
) {
	semibotroute.SemiBotBnFuture(
		app,
		config,
		orderType,
		positionSide,
		side,
		svcRepository,
	)
}

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	lambdaroute.UpdateAWSAppConfig(app, config)
}

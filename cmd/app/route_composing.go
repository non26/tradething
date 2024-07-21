package app

import (
	route "tradething/cmd/app/route/app"
	routetimeinterval "tradething/cmd/app/route/bot_semi_mannual/timeinterval"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func RouteRestApiConposing(app *echo.Echo, config *config.AppConfig) {
	route.BkRouting(
		app,
		config.Secrets.BitkubApikey,
		config.Secrets.BitkubSecretKey,
		config.ServiceName.KubSpot,
		&config.KubSpotUrl,
	)
	route.BnRouting(
		app,
		config.ServiceName.BinanceFuture,
		&config.Secrets,
		&config.BinanceFutureUrl)
	route.KcRouting(
		app,
		config.ServiceName.KucoinSpot,
		config.ServiceName.KucoinFuture,
		config.Secrets.KucoinApiKey,
		config.Secrets.KucoinApiKeyVersion,
		config.Secrets.KucoinSecretKey,
		config.Secrets.KucoinPassphase,
		&config.KCFutureUrl,
		&config.KCSpotUrl,
	)
	route.OkxRouting(
		app,
		config.Secrets.OkxApiKey,
		config.Secrets.OkxSecretKey,
		config.Secrets.OkxPassPhase,
		config.ServiceName.OKXFuture,
		&config.OkxFutureUrl,
		&config.Secrets,
		config.Env,
	)
}

func RouteSemiBotRestApiConposing(app *echo.Echo, config *config.AppConfig) {
	routetimeinterval.TimeIntervalRoute(
		app,
		config,
	)
}
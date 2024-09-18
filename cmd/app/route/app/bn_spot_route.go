package route

import (
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func BnSpotRouting(
	app *echo.Echo,
	serviceName string,
	secret *config.Secrets,
	bnFutureConfig *config.BinanceFutureUrl,
) {
}

package route

import (
	bnservice "tradething/app/bn/bn_future/bnservice"
	bothandler "tradething/app/bn/bn_future/bot_handler"
	botservice "tradething/app/bn/bn_future/bot_service"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func SemiBotBnFuture(
	app *echo.Echo,
	config *config.AppConfig,
) {
	service_name := "bn-future-semibot"

	bnservice := bnservice.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
	)
	_botservice := botservice.NewBotService(bnservice)

	_bothandler := bothandler.NewBotHandler(config, service_name, _botservice)
	bnTimeIntervalGroup := app.Group("/" + service_name)
	bnTimeIntervalGroup.POST("/time-frame-interval", _bothandler.BotTimeFrameIntervalHandler)
}

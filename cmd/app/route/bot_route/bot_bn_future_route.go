package route

import (
	bnservice "tradething/app/bn/bn_future/bnservice"
	bothandler "tradething/app/bn/bn_future/bot_handler"
	botservice "tradething/app/bn/bn_future/bot_service"
	svcrepository "tradething/app/bn/bn_future/repository"
	"tradething/app/bn/bncommon"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func SemiBotBnFuture(
	app *echo.Echo,
	config *config.AppConfig,
	orderType bncommon.IOrderType,
	positionSide bncommon.IPositionSide,
	side bncommon.ISide,
	svcRepository svcrepository.IRepository,
) {
	service_name := "bn-future-semibot"

	bnservice := bnservice.NewBinanceFutureExternalService(
		&config.BinanceFutureUrl,
		&config.Secrets,
		config.ServiceName.BinanceFuture,
	)
	botservice := botservice.NewBotService(bnservice, svcRepository, orderType, positionSide, side)

	bothandler := bothandler.NewBotHandler(config, service_name, botservice)
	bnTimeIntervalGroup := app.Group("/" + service_name)
	bnTimeIntervalGroup.POST("/time-frame-interval", bothandler.BotTimeFrameIntervalHandler)
}

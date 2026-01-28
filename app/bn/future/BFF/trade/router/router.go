package router

import (
	"tradething/app/bn/future/BFF/trade/handler"
	tradeadaptor "tradething/app/bn/future/BFF/trade/infrastructure/adaptor/trade"
	accumapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/accumulation"
	positionapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position"
	positionhistoryapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/position_history"
	subaccountapi "tradething/app/bn/future/BFF/trade/infrastructure/externalapi/sub_account"

	"tradething/app/bn/future/BFF/trade/service"
	accumulationcoreservice "tradething/app/bn/future/accumulation/service"
	positioncoreservice "tradething/app/bn/future/position/service"
	positionhistorycoreservice "tradething/app/bn/future/position_history/service"
	subaccountcoreservice "tradething/app/bn/future/sub_account/service"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func Router(
	app *echo.Echo,
	config *config.AppConfig,
	accumulationCoreService accumulationcoreservice.IService,
	cuurentPositonCoreService positioncoreservice.ICurrentPositionService,
	advancedPositionCoreService positioncoreservice.IAdvancedPositionService,
	positionHistoryCoreService positionhistorycoreservice.IService,
	subaccountCoreService subaccountcoreservice.ISubAccountService,
) {

	subaccountExternalService := subaccountapi.NewSubAccountExternalService(subaccountCoreService)
	positionHistoryExternalServce := positionhistoryapi.NewPositionHistoryExternalService(positionHistoryCoreService)
	adanvanedPositionExternalService := positionapi.NewAdvancedPositionExternalService(advancedPositionCoreService)
	currentPositionExternalService := positionapi.NewCurrentPositionExternalService(cuurentPositonCoreService)
	accumExternalService := accumapi.NewAccumulationExternalService(accumulationCoreService)
	tradeAdaptor := tradeadaptor.NewTradeAdaptor(config.BinanceAdaptorFutureUsdt.BaseUrl, config.BinanceAdaptorFutureUsdt.NewOrder)

	tradeService := service.NewTradeService(
		tradeAdaptor,
		accumExternalService,
		currentPositionExternalService,
		adanvanedPositionExternalService,
		positionHistoryExternalServce,
		subaccountExternalService,
	)

	newOrderHandler := handler.NewNewOrderHandler(tradeService)
	closeOrderByIdHandler := handler.NewCloseOrderByIdHandler(tradeService)
	accumulateOrderHandler := handler.NewAccumulateOrderHandler(tradeService)

	router := app.Group("/trade")
	router.POST("/new-order", newOrderHandler.Handle)
	router.POST("/close-order-by-id", closeOrderByIdHandler.Handle)
	router.POST("/accumulate-order", accumulateOrderHandler.Handle)

}

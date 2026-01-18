package route

import (
	"tradething/app/bn/future/BFF/market_data/handler"
	infraexternalservice "tradething/app/bn/future/BFF/market_data/infrastructure/external_service/market_data"
	"tradething/app/bn/future/BFF/market_data/service"

	externalservice "tradething/app/bn/future/market_data/service"

	"github.com/labstack/echo/v4"
)

func Router(app *echo.Echo, ext externalservice.IService) {
	router := app.Group("/market-data")

	infra := infraexternalservice.NewMarketData(ext)
	service := service.NewService(infra)

	getKlineHandler := handler.NewGetKlineHandler(service)
	router.POST("/kline", getKlineHandler.Handler)

	getPreviousKlineHandler := handler.NewGetPreviousKlineHandler(service)
	router.POST("/previous/kline", getPreviousKlineHandler.Handler)

}

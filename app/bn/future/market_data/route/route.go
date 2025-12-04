package route

import (
	"tradething/app/bn/future/market_data/handler"
	adaptor "tradething/app/bn/future/market_data/infrastructure/adaptor/market_data"
	"tradething/app/bn/future/market_data/service"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, config *config.AppConfig) {

	marketDataAdaptor := adaptor.NewMarketDataAdaptor(
		config.BinanceAdaptorFutureUsdt.BaseUrl,
		config.BinanceAdaptorFutureUsdt.KlinesCandleStick,
	)

	service := service.NewService(marketDataAdaptor)

	group := e.Group("/market-data")

	getKlineHandler := handler.NewGetKlineHandler(service)
	group.POST("/kline", getKlineHandler.Handler)

	getPreviousKlineHandler := handler.NewGetPreviousKlineHandler(service)
	group.POST("/previous/kline", getPreviousKlineHandler.Handler)

}

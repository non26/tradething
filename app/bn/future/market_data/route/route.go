package route

import (
	"tradething/app/bn/future/market_data/handler"
	adaptor "tradething/app/bn/future/market_data/infrastructure/adaptor/market_data"
	"tradething/app/bn/future/market_data/service"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo, config *config.AppConfig) {

	marketDataAdaptor := adaptor.NewMarketDataAdaptor(
		config.BinanceAdaptorFutureUsdt.BaseUrl,
		config.BinanceAdaptorFutureUsdt.KlinesCandleStick,
	)

	service := service.NewService(marketDataAdaptor)

	getKlineHandler := handler.NewGetKlineHandler(service)
	e.POST("/market-data/kline", getKlineHandler.Handler)

	getPreviousKlineHandler := handler.NewGetPreviousKlineHandler(service)
	e.POST("/market-data/previous/kline", getPreviousKlineHandler.Handler)

}

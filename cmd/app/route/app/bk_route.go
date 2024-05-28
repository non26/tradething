package route

import (
	bkservice "tradetoolv2/app/bk/app/bkservice/order"
	"tradetoolv2/app/bk/app/handler"
	service "tradetoolv2/app/bk/app/service/spot/order"
	"tradetoolv2/config"

	"github.com/labstack/echo/v4"
)

func BkRouting(
	app *echo.Echo,
	bkApi string,
	bkSecret string,
	serviceName string,
	bkConfig *config.KubSpotUrl,
) {
	bitkubGroup := app.Group("/" + serviceName)
	bitkubService := bkservice.NewOrderBkService(
		bkApi,
		bkSecret,
		serviceName,
		bkConfig,
	)
	service := service.NewOrderService(
		bitkubService,
	)
	buyOrderHandler := handler.NewBuyOrderHandler(
		service,
	)
	bitkubGroup.POST("/buy-order", buyOrderHandler.Handler)

	sellOrderHandler := handler.NewSellOrderHandler(
		service,
	)
	bitkubGroup.POST("/sellorder", sellOrderHandler.Handler)

}

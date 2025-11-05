package route

import (
	"tradething/app/bn/future/position_history/handler"
	"tradething/app/bn/future/position_history/handler/req"
	db "tradething/app/bn/future/position_history/infrastructure/db/history"
	"tradething/app/bn/future/position_history/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func Router(app *echo.Echo, dbclient *dynamodb.Client) {
	historyRepository := db.NewBnFtHistoryRepository(dbclient)
	service := service.NewService(historyRepository)

	insertHandler := handler.NewInsertHistoryHandler[req.InsertReq](service)
	app.POST("/position-history/insert", insertHandler.Handler)

	getHandler := handler.NewGetHistoryHandler[string](service)
	app.GET("/position-history/:clientId", getHandler.Handler)
}

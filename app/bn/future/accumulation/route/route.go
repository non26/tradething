package route

import (
	"tradething/app/bn/future/accumulation/handler"
	db "tradething/app/bn/future/accumulation/infrastructure/db/accumulation"
	"tradething/app/bn/future/accumulation/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo, dbclient *dynamodb.Client) {
	group := e.Group("/accumulation")

	repository := db.NewBnFtAccumulationRepository(dbclient)
	service := service.NewBnFtAccumulationService(repository)

	upsertHandler := handler.NewUpsertHandler(service)
	group.POST("/upsert", upsertHandler.Handler)

	getHandler := handler.NewGetHandler(service)
	group.POST("/get", getHandler.Handler)

	getAllHandler := handler.NewGetAllHandler(service)
	group.GET("/all", getAllHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(service)
	group.DELETE("/delete", deleteHandler.Handler)
}

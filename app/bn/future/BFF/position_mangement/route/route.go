package route

import (
	"tradething/app/bn/future/BFF/position_mangement/handler"
	externalapi "tradething/app/bn/future/BFF/position_mangement/infrastructure/external_api"
	"tradething/app/bn/future/BFF/position_mangement/service"
	positionService "tradething/app/bn/future/position/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func Router(
	app *echo.Echo,
	dbclient *dynamodb.Client,
	currentPositionService positionService.ICurrentPositionService,
	advancedPositionService positionService.IAdvancedPositionService,
) {
	router := app.Group("/position-mangement")

	externalService := externalapi.NewPositionService(currentPositionService, advancedPositionService)
	service := service.NewService(externalService)

	getHandler := handler.NewGetHandler(service)
	router.POST("/get", getHandler.Handler)

	getAllHandler := handler.NewGetAllHandler(service)
	router.GET("/all", getAllHandler.Handler)

	upsertHandler := handler.NewUpsertHandler(service)
	router.POST("/upsert", upsertHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(service)
	router.DELETE("/delete", deleteHandler.Handler)
}

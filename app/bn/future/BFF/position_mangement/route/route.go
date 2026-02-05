package route

import (
	"tradething/app/bn/future/BFF/position_mangement/handler"
	externalapi "tradething/app/bn/future/BFF/position_mangement/infrastructure/external_api"
	"tradething/app/bn/future/BFF/position_mangement/service"
	positionService "tradething/app/bn/future/position/service"
	positionhistoryservice "tradething/app/bn/future/position_history/service"

	"github.com/labstack/echo/v4"
)

func Router(
	app *echo.Group,
	currentPositionService positionService.ICurrentPositionService,
	positionHistoryService positionhistoryservice.IService,
) {
	router := app.Group("/position-mangement")

	positionRepository := externalapi.NewHistoryPositionExternalService(positionHistoryService)
	externalService := externalapi.NewPositionService(currentPositionService)
	service := service.NewService(externalService, positionRepository)

	getHandler := handler.NewGetHandler(service)
	router.POST("/get", getHandler.Handler)

	getAllHandler := handler.NewGetAllHandler(service)
	router.GET("/all", getAllHandler.Handler)

	upsertHandler := handler.NewUpsertHandler(service)
	router.POST("/upsert", upsertHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(service)
	router.POST("/delete", deleteHandler.Handler)
}

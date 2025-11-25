package route

import (
	"tradething/app/bn/future/BFF/advanced_position/handler"
	"tradething/app/bn/future/BFF/advanced_position/infrastructure/externalapi"
	"tradething/app/bn/future/BFF/advanced_position/service"
	advPositionService "tradething/app/bn/future/position/service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, advPositionService advPositionService.IAdvancedPositionService) {
	router := e.Group("/advanced-position")

	externalService := externalapi.NewAdvancedPositionExternalService(advPositionService)

	service := service.NewAdvancedPositionService(externalService)

	insertHandler := handler.NewInsertHandler(service)
	router.POST("/insert", insertHandler.Handler)

	updateHandler := handler.NewUpdateHandler(service)
	router.PUT("/update", updateHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(service)
	router.DELETE("/delete", deleteHandler.Handler)

	getHandler := handler.NewGetHandler(service)
	router.GET("/get", getHandler.Handler)

	getAllHandler := handler.NewGetAllHandler(service)
	router.GET("/all", getAllHandler.Handler)
}

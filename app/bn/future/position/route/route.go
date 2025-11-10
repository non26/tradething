package route

import (
	handleradv "tradething/app/bn/future/position/handler/advanced_position"
	handlercurrent "tradething/app/bn/future/position/handler/current_position"
	dbadvanced "tradething/app/bn/future/position/infrastructure/db/adanced_position"
	dbcurrent "tradething/app/bn/future/position/infrastructure/db/opening_position"
	serviceadv "tradething/app/bn/future/position/service/advanced_position"
	servicecurrent "tradething/app/bn/future/position/service/current_position"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func Router(app *echo.Echo, dbclient *dynamodb.Client) {
	routerCurrentPosition(app, dbclient)
	routerAdvancedPosition(app, dbclient)

}

func routerCurrentPosition(app *echo.Echo, dbclient *dynamodb.Client) {
	currentPositionRepository := dbcurrent.NewOpeningPositionRepository(dbclient)
	currentPositionService := servicecurrent.NewCurrentPositionService(currentPositionRepository)
	getHandler := handlercurrent.NewGetHandler(currentPositionService)
	getAllHandler := handlercurrent.NewGetAllHandler(currentPositionService)
	upsertHandler := handlercurrent.NewUpsertHandler(currentPositionService)
	deleteHandler := handlercurrent.NewDeleteHandler(currentPositionService)

	app.POST("/position", getHandler.Handler)
	app.POST("/position/all", getAllHandler.Handler)
	app.POST("/position/upsert", upsertHandler.Handler)
	app.POST("/position/delete", deleteHandler.Handler)
}

func routerAdvancedPosition(app *echo.Echo, dbclient *dynamodb.Client) {
	advancedPositionRepository := dbadvanced.NewAdvancedPositionRepository(dbclient)
	advancedPositionService := serviceadv.NewAdvancedPositionService(advancedPositionRepository)
	getHandler := handleradv.NewGetHandler(advancedPositionService)
	getAllHandler := handleradv.NewGetAllHandler(advancedPositionService)
	upsertHandler := handleradv.NewUpsertHandler(advancedPositionService)
	deleteHandler := handleradv.NewDeleteHandler(advancedPositionService)

	app.POST("/advanced-position", getHandler.Handler)
	app.POST("/advanced-position/all", getAllHandler.Handler)
	app.POST("/advanced-position/upsert", upsertHandler.Handler)
	app.POST("/advanced-position/delete", deleteHandler.Handler)
}

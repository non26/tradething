package route

import (
	"tradething/app/bn/future/sub_account/handler"
	"tradething/app/bn/future/sub_account/handler/req"
	db "tradething/app/bn/future/sub_account/infrastructure/db/register_account"
	"tradething/app/bn/future/sub_account/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func Router(app *echo.Echo, dbclient *dynamodb.Client) {

	subAccountRepository := db.NewRegisterAccountRepository(dbclient)
	service := service.NewSubAccountService(subAccountRepository)
	getHandler := handler.NewGetHandler[string](service)
	getAllHandler := handler.NewGetAllHandler[any](service)
	upsertHandler := handler.NewUpsertHandler[req.UpsertReq](service)
	deleteHandler := handler.NewDeleteHandler[req.DeleteReq](service)

	app.GET("/sub-account/:accountId", getHandler.Handler)
	app.GET("/sub-account/all", getAllHandler.Handler)
	app.POST("/sub-account/upsert", upsertHandler.Handler)
	app.POST("/sub-account/delete", deleteHandler.Handler)
}

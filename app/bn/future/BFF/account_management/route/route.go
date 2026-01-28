package route

import (
	"tradething/app/bn/future/BFF/account_management/handler"
	"tradething/app/bn/future/BFF/account_management/infrastructure/externalapi"
	"tradething/app/bn/future/BFF/account_management/service"
	subaccountService "tradething/app/bn/future/sub_account/service"

	"github.com/labstack/echo/v4"
)

func Router(app *echo.Echo, subaccountService subaccountService.ISubAccountService) {
	group := app.Group("/account-management")

	subAccountExternalService := externalapi.NewSubAccountExternalService(subaccountService)
	accountManagementService := service.NewAccountManagementService(subAccountExternalService)

	getAllHandler := handler.NewGetAllHandler(accountManagementService)
	group.GET("/all", getAllHandler.Handler)

	insertHandler := handler.NewInsertHandler(accountManagementService)
	group.POST("/insert", insertHandler.Handler)

	updateHandler := handler.NewUpdateHandler(accountManagementService)
	group.POST("/update", updateHandler.Handler)

	deleteHandler := handler.NewDeleteHandler(accountManagementService)
	group.POST("/delete", deleteHandler.Handler)
}

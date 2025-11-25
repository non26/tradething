package handler

import (
	"net/http"
	"tradething/app/bn/future/BFF/account_management/handler/res"
	"tradething/app/bn/future/BFF/account_management/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getAllHandler struct {
	service service.IAccountManagementService
}

func NewGetAllHandler(service service.IAccountManagementService) *getAllHandler {
	return &getAllHandler{service: service}
}

func (h *getAllHandler) Handler(c echo.Context) error {
	accounts, err := h.service.GetAllSubAccount(c.Request().Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	res := res.GetAllRes{}
	res.FromDomain(accounts)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, res)
	return response.SendResponse(http.StatusOK, c)
}

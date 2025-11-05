package handler

import (
	"net/http"
	"tradething/app/bn/future/sub_account/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getallHandler[Req any] struct {
	service service.ISubAccountService
}

func NewGetAllHandler[Req any](service service.ISubAccountService) IHandler[Req] {
	return &getallHandler[Req]{service: service}
}

func (h *getallHandler[Req]) GetReqBody(c echo.Context) (*Req, error) {
	return nil, nil
}

func (h *getallHandler[Req]) Handler(c echo.Context) error {
	subAccounts, err := h.service.GetAllSubAccount(c.Request().Context())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, subAccounts)
	return response.SendResponse(http.StatusOK, c)
}

package handler

import (
	"errors"
	"net/http"
	"tradething/app/bn/future/sub_account/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHandler[Req any] struct {
	service service.ISubAccountService
}

func NewGetHandler[Req any](service service.ISubAccountService) IHandler[string] {
	return &getHandler[string]{service: service}
}

func (h *getHandler[Req]) GetReqBody(c echo.Context) (*string, error) {
	accountId := c.Param("accountId")
	if accountId == "" {
		return nil, errors.New("accountId is required")
	}
	return &accountId, nil
}

func (h *getHandler[Req]) Handler(c echo.Context) error {

	accountId, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	subAccount, err := h.service.GetSubAccount(c.Request().Context(), *accountId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, subAccount)
	return response.SendResponse(http.StatusOK, c)
}

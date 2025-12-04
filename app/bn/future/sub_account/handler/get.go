package handler

import (
	"net/http"
	"tradething/app/bn/future/sub_account/handler/req"
	"tradething/app/bn/future/sub_account/handler/res"
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
	req := &req.GetReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return &req.AccountId, nil
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

	responseData := &res.GetSubAccountRes{}
	responseData = responseData.FromDomain(subAccount)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	return response.SendResponse(http.StatusOK, c)
}

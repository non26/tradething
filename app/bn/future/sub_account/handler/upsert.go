package handler

import (
	"net/http"
	"tradething/app/bn/future/sub_account/handler/req"
	"tradething/app/bn/future/sub_account/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type upsertHandler[Req any] struct {
	service service.ISubAccountService
}

func NewUpsertHandler[Req any](service service.ISubAccountService) IHandler[req.UpsertReq] {
	return &upsertHandler[req.UpsertReq]{service: service}
}

func (h *upsertHandler[Req]) GetReqBody(c echo.Context) (*req.UpsertReq, error) {
	req := &req.UpsertReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *upsertHandler[Req]) Handler(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	if err := reqBody.Validate(); err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.service.UpsertSubAccount(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

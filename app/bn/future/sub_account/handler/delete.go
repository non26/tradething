package handler

import (
	"net/http"
	"tradething/app/bn/future/sub_account/handler/req"
	"tradething/app/bn/future/sub_account/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type deleteHandler[Req any] struct {
	service service.ISubAccountService
}

func NewDeleteHandler[Req any](service service.ISubAccountService) IHandler[req.DeleteReq] {
	return &deleteHandler[req.DeleteReq]{service: service}
}

func (h *deleteHandler[Req]) GetReqBody(c echo.Context) (*req.DeleteReq, error) {
	req := &req.DeleteReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *deleteHandler[Req]) Handler(c echo.Context) error {

	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.service.DeleteSubAccount(c.Request().Context(), reqBody.AccountId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

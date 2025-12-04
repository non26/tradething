package handler

import (
	"net/http"
	"tradething/app/bn/future/position_history/handler/req"
	"tradething/app/bn/future/position_history/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type insertHistoryHandler[Req any] struct {
	service service.IService
}

func NewInsertHistoryHandler[Req any](service service.IService) IHandler[req.InsertReq] {
	return &insertHistoryHandler[req.InsertReq]{service: service}
}

func (h *insertHistoryHandler[Req]) GetReqBody(c echo.Context) (*req.InsertReq, error) {
	req := &req.InsertReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *insertHistoryHandler[Req]) Handler(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.service.InsertPositionHistory(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

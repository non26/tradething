package handler

import (
	"errors"
	"net/http"
	"tradething/app/bn/future/position_history/handler/res"
	"tradething/app/bn/future/position_history/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getHistoryHandler[Req any] struct {
	service service.IService
}

func NewGetHistoryHandler[Req any](service service.IService) IHandler[string] {
	return &getHistoryHandler[string]{service: service}
}

func (h *getHistoryHandler[Req]) GetReqBody(c echo.Context) (*string, error) {
	clientId := c.Param("clientId")
	if clientId == "" {
		return nil, errors.New("clientId is required")
	}
	return &clientId, nil
}

func (h *getHistoryHandler[Req]) Handler(c echo.Context) error {
	clientId, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	history, err := h.service.GetPositionHistory(c.Request().Context(), *clientId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	responseData := &res.GetHistoryRes{}
	responseData = responseData.FromDomain(history)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, responseData)
	return response.SendResponse(http.StatusOK, c)
}

package handler

import (
	"net/http"
	"tradething/app/bn/future/market_data/handler/req"
	"tradething/app/bn/future/market_data/handler/res"
	"tradething/app/bn/future/market_data/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getPreviousKlineHandler struct {
	service service.IService
}

func NewGetPreviousKlineHandler(service service.IService) *getPreviousKlineHandler {
	return &getPreviousKlineHandler{service: service}
}

func (h *getPreviousKlineHandler) Handler(c echo.Context) error {
	req := new(req.GetKlineReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	kline, err := h.service.GetPreviousKline(c.Request().Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	klineResponse := &res.KlineResponse{}
	klineResponse = klineResponse.FromDomain(kline)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, klineResponse)
	return response.SendResponse(http.StatusOK, c)
}

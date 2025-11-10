package handler

import (
	"net/http"
	"tradething/app/bn/future/market_data/handler/req"
	"tradething/app/bn/future/market_data/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type getKlineHandler struct {
	service service.IService
}

func NewGetKlineHandler(service service.IService) *getKlineHandler {
	return &getKlineHandler{service: service}
}

func (h *getKlineHandler) Handler(c echo.Context) error {
	req := new(req.GetKlineReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	kline, err := h.service.GetKline(c.Request().Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, kline)
	return response.SendResponse(http.StatusOK, c)
}

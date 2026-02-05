package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tradething/app/bn/future/BFF/market_data/handler/req"
	"tradething/app/bn/future/BFF/market_data/handler/res"
	"tradething/app/bn/future/BFF/market_data/service"

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
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}
	stringifyReq, _ := json.Marshal(req)
	fmt.Print("req from handler", string(stringifyReq))

	klines, err := h.service.GetKline(c.Request().Context(), req.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	klinesResponse := &res.KlinesResponse{}
	klinesResponse = klinesResponse.FromDomain(klines)
	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, klinesResponse)
	return response.SendResponse(http.StatusOK, c)
}

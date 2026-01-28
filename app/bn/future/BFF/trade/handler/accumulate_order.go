package handler

import (
	"net/http"
	"tradething/app/bn/future/BFF/trade/handler/req"
	"tradething/app/bn/future/BFF/trade/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type accumulateOrderHandler struct {
	tradeService service.ITradeService
}

func NewAccumulateOrderHandler(tradeService service.ITradeService) *accumulateOrderHandler {
	return &accumulateOrderHandler{tradeService: tradeService}
}

func (h *accumulateOrderHandler) GetReqBody(c echo.Context) (*req.AccumulateOrderReq, error) {
	req := &req.AccumulateOrderReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *accumulateOrderHandler) Handle(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.tradeService.AccumulateOrder(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

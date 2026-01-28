package handler

import (
	"net/http"
	"tradething/app/bn/future/BFF/trade/handler/req"
	"tradething/app/bn/future/BFF/trade/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type newOrderHandler struct {
	tradeService service.ITradeService
}

func NewNewOrderHandler(tradeService service.ITradeService) *newOrderHandler {
	return &newOrderHandler{tradeService: tradeService}
}

func (h *newOrderHandler) GetReqBody(c echo.Context) (*req.NewOrderReq, error) {
	req := &req.NewOrderReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *newOrderHandler) Handle(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.tradeService.NewOrder(c.Request().Context(), reqBody.ToDomain())
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

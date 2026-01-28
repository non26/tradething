package handler

import (
	"net/http"
	"tradething/app/bn/future/BFF/trade/handler/req"
	"tradething/app/bn/future/BFF/trade/service"

	"github.com/labstack/echo/v4"
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type closeOrderByIdHandler struct {
	tradeService service.ITradeService
}

func NewCloseOrderByIdHandler(tradeService service.ITradeService) *closeOrderByIdHandler {
	return &closeOrderByIdHandler{tradeService: tradeService}
}

func (h *closeOrderByIdHandler) GetReqBody(c echo.Context) (*req.CloseOrderByIdReq, error) {
	req := &req.CloseOrderByIdReq{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *closeOrderByIdHandler) Handle(c echo.Context) error {
	reqBody, err := h.GetReqBody(c)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.InvalidRequestErrorCode, err.Error(), nil)
		return response.SendResponse(http.StatusBadRequest, c)
	}

	err = h.tradeService.CloseOrderById(c.Request().Context(), reqBody.ClientId)
	if err != nil {
		response := appresponse.NewAppResponse(appresponse.FailCode, err.Error(), nil)
		return response.SendResponse(http.StatusInternalServerError, c)
	}

	response := appresponse.NewAppResponse(appresponse.SuccessCode, appresponse.SuccessMsg, nil)
	return response.SendResponse(http.StatusOK, c)
}

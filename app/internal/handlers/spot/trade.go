package handlers

import (
	"net/http"
	"tradething/app/internal/handlers/spot/req"
	process "tradething/app/internal/process/spot"

	"github.com/labstack/echo/v4"
)

type ITradeHandler interface {
	Trade(c echo.Context) error
}

type TradeHandler struct {
	spotProcess process.ISpot
}

func NewTradeHandler(
	spotProcess process.ISpot,
) ITradeHandler {
	return &TradeHandler{
		spotProcess: spotProcess,
	}
}

func (h *TradeHandler) Trade(c echo.Context) error {
	request := req.Trade{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.spotProcess.PlaceOrder(c.Request().Context(), request.ToOrder())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

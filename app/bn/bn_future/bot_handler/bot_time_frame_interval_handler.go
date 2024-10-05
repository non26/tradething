package bnfuture

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"
)

func (h *botHandler) BotTimeFrameIntervalHandler(c echo.Context) error {
	payload := new(bothandlerreq.TradeTimeIntervalBinanceFutureRequest)
	err := c.Bind(payload)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			&bothandlerres.TradeTimeIntervalBinanceFutureResponse{
				Message: err.Error(),
			},
		)
	}
	payload.ToUpper()

	data, err := h.semibot.TimeIntervalSemiBotService(
		c.Request().Context(),
		payload,
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&bothandlerres.TradeTimeIntervalBinanceFutureResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		data,
	)
}

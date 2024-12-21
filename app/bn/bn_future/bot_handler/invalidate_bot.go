package bnfuture

import (
	"net/http"
	bnftbotereq "tradething/app/bn/bn_future/bot_handler_request_model"
	bnftboteres "tradething/app/bn/bn_future/bot_handler_response_model"
	bot_service "tradething/app/bn/bn_future/bot_service"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IInvalidateBotHandler interface {
	Handler(c echo.Context) error
	GetRequestBody(c echo.Context) (*bnftbotereq.InvalidateBotHandlerRequest, error)
}

type InvalidateBotHandler struct {
	botService bot_service.IBotService
}

func NewInvalidateBotHandler(
	botService bot_service.IBotService,
) IInvalidateBotHandler {
	return &InvalidateBotHandler{
		botService: botService,
	}
}

func (h *InvalidateBotHandler) GetRequestBody(c echo.Context) (*bnftbotereq.InvalidateBotHandlerRequest, error) {
	req := new(bnftbotereq.InvalidateBotHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *InvalidateBotHandler) Handler(c echo.Context) error {
	req, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.CommonResponse{
			Code:    common.FailCode,
			Message: err.Error(),
		})
	}

	response, err := h.botService.InvalidateBot(c.Request().Context(), req.ToServiceModel())
	if err != nil {
		res := &bnftboteres.InvalidateBotHandlerResponse{
			BotId:      req.BotId,
			BotOrderId: req.BotOrderId,
			Message:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, response)
}

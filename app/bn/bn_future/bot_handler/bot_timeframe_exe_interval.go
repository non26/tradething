package bnfuture

import (
	"net/http"
	bot_service "tradething/app/bn/bn_future/bot"
	bnftbotereq "tradething/app/bn/bn_future/bot_handler_request"
	bnftboteres "tradething/app/bn/bn_future/bot_handler_response"
	"tradething/common"

	"github.com/labstack/echo/v4"
)

type IBotTimeframeExeIntervalHandler interface {
	Handler(c echo.Context) error
	GetRequestBody(c echo.Context) (*bnftbotereq.BotTimeframeExeIntervalHandlerRequest, error)
}

type botTimeframeExeIntervalHandler struct {
	botService bot_service.IBotService
}

func NewBotTimeframeExeIntervalHandler(
	botService bot_service.IBotService,
) IBotTimeframeExeIntervalHandler {
	return &botTimeframeExeIntervalHandler{
		botService: botService,
	}
}

func (h *botTimeframeExeIntervalHandler) GetRequestBody(c echo.Context) (*bnftbotereq.BotTimeframeExeIntervalHandlerRequest, error) {
	req := new(bnftbotereq.BotTimeframeExeIntervalHandlerRequest)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h *botTimeframeExeIntervalHandler) Handler(c echo.Context) error {
	req, err := h.GetRequestBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.CommonResponse{
			Code:    common.FailCode,
			Message: err.Error(),
		})
	}
	req.Transform()
	svcReq, err := req.ToBotServiceRequest()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.CommonResponse{
			Code:    common.FailCode,
			Message: err.Error(),
		})
	}
	response, err := h.botService.BotTimeframeExeInterval(c.Request().Context(), svcReq)
	if err != nil {
		res := &bnftboteres.BotTimeframeExeIntervalResponse{
			BotOrderID: req.BotOrderID,
			Message:    err.Error(),
			Status:     "failed",
			Code:       common.FailCode,
		}
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, response)
}

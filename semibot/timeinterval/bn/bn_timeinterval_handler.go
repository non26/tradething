package timeinterval

import (
	"net/http"
	"strings"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

type TradeTimeIntervalBinanceFutureRequest struct {
	// Side          string  `json:"side"`
	PositionSide    string  `json:"positionSide"`  // long/short
	EntryQuantity   float64 `json:"entryQuantity"` // 0.005
	Symbol          string  `json:"symbol"`        // btcusdt
	PrevClientId    string  `json:"prevCliId"`
	CurrentClientId string  `json:"currCliId"`
	LeverageLevel   string  `json:"leverageLevel"` // 125
}

func (t *TradeTimeIntervalBinanceFutureRequest) ToUpper() {
	t.PositionSide = strings.ToUpper(t.PositionSide)
	t.Symbol = strings.ToUpper(t.Symbol)
}

type TradeTimeIntervalBinanceFutureResponse struct {
	Message string `json:"message"`
}

type bnHandler struct {
	log          *config.AppConfig
	service_name string
	semibot      *bnTimeInterval
}

func NewBnTradeTimeIntervalHandler(
	log *config.AppConfig,
	service_name string,
	semibot *bnTimeInterval,
) *bnHandler {
	return &bnHandler{
		log,
		service_name,
		semibot,
	}
}

func (h *bnHandler) BnHandler(c echo.Context) error {
	payload := new(TradeTimeIntervalBinanceFutureRequest)
	err := c.Bind(payload)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			&TradeTimeIntervalBinanceFutureResponse{
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
			&TradeTimeIntervalBinanceFutureResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		data,
	)
}

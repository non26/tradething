package tradepertimeframe

import (
	"context"
	"net/http"
	"strconv"
	model "tradething/app/bn/app/model/handlermodel/future"
	service "tradething/app/bn/app/service/future"

	"github.com/labstack/echo/v4"
)

type ITradePerTimeFrameBinanceFuture interface {
	QueryOrder(
		ctx context.Context,
		requet *TradePerTimeFrameBinanceFutureRequest,
	) (quatity float64, err error)
	CloseOrder()
	OpenOrder()
	TrafrPerTimeFrameSemiBot()
	TrafrPerTimeFrameSemiBotHandler(c echo.Context) error
}

type tradePerTimeFrameBinanceFuture struct {
	tradeService service.IBinanceFutureService
}

func NewTradePerTimeFrameBinanceFuture(
	tradeService service.IBinanceFutureService,
) ITradePerTimeFrameBinanceFuture {
	return &tradePerTimeFrameBinanceFuture{
		tradeService,
	}
}

func (t *tradePerTimeFrameBinanceFuture) QueryOrder(
	ctx context.Context,
	requet *TradePerTimeFrameBinanceFutureRequest,
) (quatity float64, err error) {
	m, err := t.tradeService.QueryOrder(
		ctx,
		&model.QueryOrderBinanceHandlerRequest{
			Symbol: requet.Symbol,
		},
	)
	if err != nil {
		return 0, err
	}
	quatity, err = strconv.ParseFloat(m.ExecutedQty, 64)
	if err != nil {
		return 0, err
	}
	return quatity, nil
}
func (t *tradePerTimeFrameBinanceFuture) CloseOrder() {}
func (t *tradePerTimeFrameBinanceFuture) OpenOrder()  {}
func (t *tradePerTimeFrameBinanceFuture) TrafrPerTimeFrameSemiBot() {

}

func (t *tradePerTimeFrameBinanceFuture) TrafrPerTimeFrameSemiBotHandler(c echo.Context) error {
	request := new(TradePerTimeFrameBinanceFutureRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(
			http.StatusBadGateway,
			nil,
		)
	}

	return nil
}

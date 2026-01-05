package trade

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/trade/req"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/trade/res"
)

type tradeAdaptor struct {
	baseUrl          string
	newOrderEndPoint string
}

func NewTradeAdaptor(baseUrl string, newOrderEndPoint string) *tradeAdaptor {
	return &tradeAdaptor{baseUrl: baseUrl, newOrderEndPoint: newOrderEndPoint}
}

func (t *tradeAdaptor) NewOrder(ctx context.Context, order *domain.Order) error {

	method := http.MethodPost
	orderRequest := req.NewOrderReq{}
	orderRequest.ToTradeAdaptorNewOrderRequest(order)

	url := fmt.Sprintf("%s%s", t.baseUrl, t.newOrderEndPoint)
	payload, err := json.Marshal(orderRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res res.NewOrderRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.IsError() {
		return res.GetError()
	}

	return nil
}

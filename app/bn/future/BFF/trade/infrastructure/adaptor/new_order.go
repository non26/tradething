package adaptor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/req"
	"tradething/app/bn/future/BFF/trade/infrastructure/adaptor/res"
)

func (a *adaptor) NewOrder(ctx context.Context, request *req.NewOrderRequest) (*res.NewOrderResponse, error) {
	newClient := http.Client{
		Timeout: 10 * time.Second,
	}

	method := http.MethodPost
	fullUrl := fmt.Sprintf("%s%s", a.baseUrl, a.newOrderEndpoint)
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, fullUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := newClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newOrderResponse res.NewOrderResponse
	err = json.Unmarshal(body, &newOrderResponse)
	if err != nil {
		return nil, err
	}

	if newOrderResponse.Code != nil && newOrderResponse.Message != nil {
		return nil, fmt.Errorf("code: %d, message: %s", *newOrderResponse.Code, *newOrderResponse.Message)
	}

	return &newOrderResponse, nil
}

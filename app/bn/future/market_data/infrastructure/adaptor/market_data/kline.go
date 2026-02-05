package marketdata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tradething/app/bn/future/market_data/domain"
	"tradething/app/bn/future/market_data/infrastructure/adaptor/market_data/res"
)

func (a *marketDataAdaptor) GetKline(ctx context.Context, kline *domain.Kline) (*res.KlineReaderAdaptorResponse, error) {

	method := http.MethodPost
	requestBody := kline.ToMarketDataAdaptorKlineRequest()
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", a.baseUrl, a.klinesCandleStickEndpoint)
	fmt.Print("url", url)
	fmt.Print("payload", string(payload))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
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

	var klineAdaptorResponse res.KlineAdaptorResponse
	err = json.Unmarshal(body, &klineAdaptorResponse)
	if err != nil {
		return nil, err
	}

	return klineAdaptorResponse.ToReader(), nil
}

package adaptor

import (
	"context"
	"net/http"
	"tradething/config"

	req "tradething/app/internal/infrastructure/adaptor/spot/req"
	res "tradething/app/internal/infrastructure/adaptor/spot/res"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBinanceSpotTradeService interface {
	PlaceOrder(ctx context.Context, req *req.SpotOrderRequest) (*res.PlacePositionData, error)
}

type binanceSpotAdaptorService struct {
	binanceSpotUrl         *config.BinanceSpotUrl
	apikey                 string
	secretkey              string
	binanceSpotServiceName string
	httpttransport         bntransport.IBinanceServiceHttpTransport
	httpclient             bnclient.IBinanceSerivceHttpClient
}

func NewBinanceSpotAdaptorService(
	binanceSpotUrl *config.BinanceSpotUrl,
	apikey string,
	secretkey string,
	binanceSpotServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBinanceSpotTradeService {
	return &binanceSpotAdaptorService{
		binanceSpotUrl:         binanceSpotUrl,
		apikey:                 apikey,
		secretkey:              secretkey,
		binanceSpotServiceName: binanceSpotServiceName,
		httpttransport:         httpttransport,
		httpclient:             httpclient,
	}
}

func (b *binanceSpotAdaptorService) PlaceOrder(ctx context.Context, request *req.SpotOrderRequest) (*res.PlacePositionData, error) {
	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[req.SpotOrderRequest](),
		bnresponse.NewBinanceServiceHttpResponse[res.PlacePositionData](),
		b.httpttransport,
		b.httpclient,
	)

	res, err := c.CallBinance(
		req.NewPlaceSignleOrderBinanceServiceRequest(request),
		b.binanceSpotUrl.BinanceSpotBaseUrl.BianceUrl1,
		b.binanceSpotUrl.SingleOrder,
		http.MethodPost,
		b.secretkey,
		b.apikey,
		b.binanceSpotServiceName,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

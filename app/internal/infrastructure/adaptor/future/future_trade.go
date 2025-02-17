package adaptor

import (
	"context"
	"net/http"

	req "tradething/app/internal/infrastructure/adaptor/future/req/future_trade"
	res "tradething/app/internal/infrastructure/adaptor/future/res/future_trade"
	"tradething/config"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBinanceFutureTradeService interface {
	PlaceOrder(
		ctx context.Context,
		request *req.PlacePosition,
	) (*res.PlacePositionData, error)
}

type binanceFutureAdaptorService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	apikey                   string
	secretkey                string
	binanceFutureServiceName string
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBinanceFutureAdaptorService() IBinanceFutureTradeService {
	return &binanceFutureAdaptorService{}
}

func (b *binanceFutureAdaptorService) PlaceOrder(
	ctx context.Context,
	request *req.PlacePosition,
) (*res.PlacePositionData, error) {

	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[req.PlacePosition](),
		bnresponse.NewBinanceServiceHttpResponse[res.PlacePositionData](),
		b.httpttransport,
		b.httpclient,
	)

	res, err := c.CallBinance(
		req.NewPlaceSignleOrderBinanceServiceRequest(request),
		b.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		b.binanceFutureUrl.SingleOrder,
		http.MethodPost,
		b.secretkey,
		b.apikey,
		b.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

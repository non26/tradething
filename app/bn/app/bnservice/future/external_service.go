package bnservice

import (
	"context"
	"net/http"
	model "tradething/app/bn/app/model/handlermodel/future"
	"tradething/config"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.SetLeverageHandlerRequest) (*http.Response, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *model.PlaceSignleOrderHandlerRequest,
	) (*http.Response, error)
	QueryOrder(
		ctx context.Context,
		request *model.QueryOrderBinanceHandlerRequest,
	) (*http.Response, error)
}

type binanceFutureExternalService struct {
	binanceFutureUrl *config.BinanceFutureUrl
	secrets          *config.Secrets
}

func NewBinanceFutureExternalService(
	binanceFutureUrl *config.BinanceFutureUrl,
	secrets *config.Secrets,
) IBinanceFutureExternalService {
	return &binanceFutureExternalService{
		binanceFutureUrl,
		secrets,
	}
}

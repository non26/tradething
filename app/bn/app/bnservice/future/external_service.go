package bnservice

import (
	"context"
	"net/http"
	model "tradetoolv2/app/bn/app/model/handlermodel/future"
	"tradetoolv2/config"
)

type IBinanceFutureExternalService interface {
	SetNewLeverage(
		ctx context.Context,
		request *model.SetLeverageHandlerRequest) (*http.Response, error)
	PlaceSingleOrder(
		ctx context.Context,
		request *model.PlaceSignleOrderHandlerRequest,
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

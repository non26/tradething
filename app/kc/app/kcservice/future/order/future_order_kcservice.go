package kcservice

import (
	"context"
	model "tradething/app/kc/app/model/kcservicemodel/future/order"
	"tradething/config"
)

type IFutureKcService interface {
	PlaceFutureOrderKcService(
		ctx context.Context,
		req *model.PlaceFutureOrderKcServiceRequest,
	) (*model.PlaceFutureOrderKcServiceResponse, error)
}

type futureKcService struct {
	apiKey      string
	apiVersion  string
	secretKey   string
	passphass   string
	serviceName string
	kcConfig    *config.KCFutureUrl
}

func NewFutureKcService(
	apiKey string,
	apiVersion string,
	secretKey string,
	passphass string,
	serviceName string,
	kcConfig *config.KCFutureUrl,
) IFutureKcService {
	return &futureKcService{
		apiKey:      apiKey,
		apiVersion:  apiVersion,
		secretKey:   secretKey,
		passphass:   passphass,
		serviceName: serviceName,
		kcConfig:    kcConfig,
	}
}

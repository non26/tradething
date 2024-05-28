package kcservice

import (
	"context"
	model "tradetoolv2/app/kc/app/model/kcservicemodel/spot/order"
	"tradetoolv2/config"
)

type ISpotKcService interface {
	PlaceSpotOrderKcService(
		ctx context.Context,
		req *model.PlaceSpotOrderKcServiceRequest,
	) (*model.PlaceSpotOrderKcServiceResponse, error)
}

type spotKcService struct {
	apiKey      string
	apiVersion  string
	secretKey   string
	passphass   string
	serviceName string
	kcConfig    *config.KCSpotUrl
}

func NewSpotKcService(
	apiKey string,
	apiVersion string,
	secretKey string,
	passphass string,
	serviceName string,
	kcConfig *config.KCSpotUrl,
) ISpotKcService {
	return &spotKcService{
		apiKey:      apiKey,
		apiVersion:  apiVersion,
		secretKey:   secretKey,
		passphass:   passphass,
		serviceName: serviceName,
		kcConfig:    kcConfig,
	}
}

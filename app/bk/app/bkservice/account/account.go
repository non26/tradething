package bkservice

import (
	"context"
	model "tradetoolv2/app/bk/app/model/bitkubservicemodel/account"
	"tradetoolv2/config"
)

type IAccountBkService interface {
	GetBalances(
		ctx context.Context,
	) (*model.BalanceBkServiceResponse, error)
}

type accountBkService struct {
	bkApiKey     string
	bkSecretKey  string
	serviceName  string
	kubConfigUrl *config.KubSpotUrl
}

func NewAccountBkService(
	bkApiKey string,
	bkSecretKey string,
	serviceName string,
	kubConfigUrl *config.KubSpotUrl,
) IAccountBkService {
	return &accountBkService{
		bkApiKey:     bkApiKey,
		bkSecretKey:  bkSecretKey,
		serviceName:  serviceName,
		kubConfigUrl: kubConfigUrl,
	}
}

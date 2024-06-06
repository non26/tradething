package bkservice

import (
	"context"
	model "tradething/app/bk/app/model/bitkubservicemodel/account"
	"tradething/config"
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

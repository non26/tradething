package bkservice

import (
	"context"
	model "tradething/app/bk/app/model/bitkubservicemodel"
	"tradething/config"
)

type IOrderBkService interface {
	SellOrder(
		ctx context.Context,
		req *model.SellBkServiceRequest) (
		*model.SellOrderBkServiceResponse,
		error)
	BuyOrder(
		ctx context.Context,
		req *model.BuyOrderBkServiceRequest) (
		*model.BuyOrderBkServiceResponse,
		error)
}

type orderBkService struct {
	bkApiKey     string
	bkSecretKey  string
	serviceName  string
	kubConfigUrl *config.KubSpotUrl
}

func NewOrderBkService(
	bkApiKey string,
	bkSecretKey string,
	serviceName string,
	kubConfigUrl *config.KubSpotUrl,
) IOrderBkService {
	return &orderBkService{
		bkApiKey:     bkApiKey,
		bkSecretKey:  bkSecretKey,
		serviceName:  serviceName,
		kubConfigUrl: kubConfigUrl,
	}
}

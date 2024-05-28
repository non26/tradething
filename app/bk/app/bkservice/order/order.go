package bkservice

import (
	"context"
	model "tradetoolv2/app/bk/app/model/bitkubservicemodel"
	"tradetoolv2/config"
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

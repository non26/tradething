package service

import (
	"context"
	"log"
	svcFuture "tradething/app/bn/bn_future/service_model"

	bntradereq "tradething/app/bn/bn_future/bnservice_response_model/trade"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func (b *binanceFutureService) closePosition(ctx context.Context, request *svcFuture.PlaceSignleOrderServiceRequest) (*bntradereq.PlaceSignleOrderBinanceServiceResponse, error) {
	closePosition, err := b.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		log.Println("error place sell order", err.Error())
		return nil, err
	}
	err = b.repository.DeleteOpenOrderBySymbolAndPositionSide(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		log.Println("error delete open order by key", err.Error())
	}
	err = b.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
	if err != nil {
		log.Println("error insert history", err.Error())
	}
	return closePosition, nil

}

func (b *binanceFutureService) openPosition(ctx context.Context, request *svcFuture.PlaceSignleOrderServiceRequest, dbQUsdt *dynamodbmodel.BnFtQouteUSDT) (*bntradereq.PlaceSignleOrderBinanceServiceResponse, error) {

	openPosition, err := b.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		log.Println("error place order", err.Error())
		return nil, err
	}

	if request.GetPositionSide() == b.positionSideType.Short() {
		dbQUsdt.SetCountingShort(dbQUsdt.GetCountingShort() + 1)
	} else {
		dbQUsdt.SetCountingLong(dbQUsdt.GetCountingLong() + 1)
	}
	err = b.repository.UpdateQouteUSDT(ctx, dbQUsdt)
	if err != nil {
		log.Println("error update qoute usdt", err.Error())
	}
	err = b.repository.InsertNewOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		log.Println("error new open order", err.Error())
	}
	return openPosition, nil
}

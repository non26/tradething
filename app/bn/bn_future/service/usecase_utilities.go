package service

import (
	"context"
	"errors"
	"log"
	"time"
	svcFuture "tradething/app/bn/bn_future/service_model"

	bntradereq "tradething/app/bn/bn_future/bnservice_response_model/trade"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
	utils "github.com/non26/tradepkg/pkg/bn/utils"
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
	err = b.bnFtOpeningPositionTable.Delete(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		log.Println("error delete open order by key", err.Error())
	}
	err = b.bnFtHistoryTable.Insert(ctx, request.ToBnPositionHistoryRepositoryModel())
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
	err = b.bnFtQouteUsdtTable.Update(ctx, dbQUsdt)
	if err != nil {
		log.Println("error update qoute usdt", err.Error())
	}
	err = b.bnFtOpeningPositionTable.Insert(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
	if err != nil {
		log.Println("error new open order", err.Error())
	}
	return openPosition, nil
}

func (b *binanceFutureService) setDefaultClientOrderId(request *svcFuture.PlaceSignleOrderServiceRequest, dbQUsdt *dynamodbmodel.BnFtQouteUSDT) {
	var counting int
	if request.IsLongPosition() {
		counting = dbQUsdt.GetCountingLong()
	} else {
		counting = dbQUsdt.GetCountingShort()
	}
	request.SetClientOrderId(utils.BinanceDefaultClientID(request.GetSymbol(), request.GetPositionSide(), counting))
}

func (b *binanceFutureService) accumulateOrder(ctx context.Context, request *svcFuture.PlaceSignleOrderServiceRequest, dbOpeningOrder *dynamodbmodel.BnFtOpeningPosition) (*bntradereq.PlaceSignleOrderBinanceServiceResponse, error) {
	// for accumulate order
	placeOrderRes, err := b.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		log.Println("error place order for accumulate order", err.Error())
		return nil, err
	}
	dbOpeningOrder.AddMoreAmountQ(request.GetAmountQ())
	err = b.bnFtOpeningPositionTable.Update(ctx, dbOpeningOrder)
	if err != nil {
		log.Println("error update open order for accumulate order", err.Error())
		return nil, err
	}
	return placeOrderRes, nil
}

func (b *binanceFutureService) getPreviousBnTimeStartAndEnd(request *svcFuture.PlaceSignleOrderServiceRequest) (*time.Time, *time.Time, error) {
	var prv_start, prv_end time.Time
	bnTime := utils.NewBinanceTime(time.Now())

	period, unit, err := utils.GetInterval(request.GetStopLoss().Interval)
	if err != nil {
		log.Println("error get interval for watching order", err.Error())
		return nil, nil, err
	}
	/// now support only hourly
	switch unit {
	case utils.Minute:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartMinuteAndEndMinute(period)
		if err != nil {
			log.Println("error get previous bn time start minute and end minute for watching order", err.Error())
			return nil, nil, err
		}
	case utils.Hour:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartHourAndEndHour(period)
		if err != nil {
			log.Println("error get previous bn time start hour and end hour for watching order", err.Error())
			return nil, nil, err
		}
	case utils.Day:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartDayAndEndDay(period)
		if err != nil {
			log.Println("error get previous bn time start day and end day for watching order", err.Error())
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("invalid interval")
	}
	return &prv_start, &prv_end, nil
}

package service

import (
	"context"
	"errors"
	"log"
	"time"
	svcModels "tradething/app/bn/bn_future/service_model"

	bntradereq "tradething/app/bn/bn_future/bnservice_response/trade"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"
	utils "github.com/non26/tradepkg/pkg/bn/utils"
)

func (b *binanceFutureService) closePosition(ctx context.Context, request *svcModels.Position) (*bntradereq.PlacePositionData, serviceerror.IError) {
	closePosition, err := b.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		log.Println("error place sell order", err.Error())
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
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

func (b *binanceFutureService) openPosition(ctx context.Context, request *svcModels.Position) (*bntradereq.PlacePositionData, serviceerror.IError) {
	dbQUsdt, err := b.bnFtQouteUsdtTable.Get(ctx, request.GetSymbol())
	if err != nil {
		log.Println("error get qoute usdt", err.Error())
		dbQUsdt = dynamodbmodel.NewBinanceFutureQouteUSTDTableRecord(request.GetSymbol(), request.IsLongPosition())
	}
	if !dbQUsdt.IsFound() {
		dbQUsdt = dynamodbmodel.NewBinanceFutureQouteUSTDTableRecord(request.GetSymbol(), request.IsLongPosition())
		err = b.bnFtQouteUsdtTable.Insert(ctx, dbQUsdt)
		if err != nil {
			log.Println("error insert new symbol qoute usdt", err.Error())
		}
	}
	if request.GetClientId() == "" {
		b.setDefaultClientOrderId(request, dbQUsdt)
	}

	openPosition, err := b.binanceService.PlaceSingleOrder(
		ctx,
		request.ToBinanceServiceModel(),
	)
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
	}

	if utils.IsShortPosition(request.GetPositionSide()) {
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

	err = b.bnFtAdvancedPositionTable.Delete(ctx, request.ToBnAdvancedPositionRepositoryModel())
	if err != nil {
		log.Println("error delete advanced position", err.Error())
	}

	return openPosition, nil
}

func (b *binanceFutureService) setDefaultClientOrderId(request *svcModels.Position, dbQUsdt *dynamodbmodel.BnFtQouteUSDT) {
	var counting int
	if request.IsLongPosition() {
		counting = dbQUsdt.GetCountingLong()
	} else {
		counting = dbQUsdt.GetCountingShort()
	}
	request.SetClientId(utils.BinanceDefaultClientID(request.GetSymbol(), request.GetPositionSide(), counting))
}

func (b *binanceFutureService) accumulateOrder(ctx context.Context, request *svcModels.Position, dbOpeningOrder *dynamodbmodel.BnFtOpeningPosition) (*bntradereq.PlacePositionData, serviceerror.IError) {
	placeOrderRes, err := b.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
	if err != nil {
		log.Println("error place order for accumulate order", err.Error())
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
	}

	dbOpeningOrder.AddMoreAmountB(request.GetAmountB())
	err = b.bnFtOpeningPositionTable.Update(ctx, dbOpeningOrder)
	if err != nil {
		log.Println("error update open order for accumulate order", err.Error())
	}

	err = b.bnFtHistoryTable.Insert(ctx, request.ToBnPositionHistoryRepositoryModel())
	if err != nil {
		log.Println("error insert history for accumulate order", err.Error())
	}

	return placeOrderRes, nil
}

func (b *binanceFutureService) getPreviousBnTimeStartAndEnd(request *svcModels.Position) (*time.Time, *time.Time, serviceerror.IError) {
	var prv_start, prv_end time.Time
	bnTime := utils.NewBinanceTime(time.Now())

	period, unit, err := utils.GetInterval(request.GetStopLoss().Interval)
	if err != nil {
		log.Println("error get interval for watching order", err.Error())
		return nil, nil, serviceerror.NewServiceErrorWith(serviceerror.BN_COMMON_ERROR, err)
	}

	switch unit {
	case utils.Minute:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartMinuteAndEndMinute(period)
		if err != nil {
			log.Println("error get previous bn time start minute and end minute for watching order", err.Error())
			return nil, nil, serviceerror.NewServiceErrorWith(serviceerror.BN_COMMON_ERROR, err)
		}
	case utils.Hour:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartHourAndEndHour(period)
		if err != nil {
			log.Println("error get previous bn time start hour and end hour for watching order", err.Error())
			return nil, nil, serviceerror.NewServiceErrorWith(serviceerror.BN_COMMON_ERROR, err)
		}
	case utils.Day:
		var err error
		prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartDayAndEndDay(period)
		if err != nil {
			log.Println("error get previous bn time start day and end day for watching order", err.Error())
			return nil, nil, serviceerror.NewServiceErrorWith(serviceerror.BN_COMMON_ERROR, err)
		}
	default:
		return nil, nil, serviceerror.NewServiceErrorWith(serviceerror.BN_COMMON_ERROR, errors.New("invalid interval"))
	}

	return &prv_start, &prv_end, nil
}

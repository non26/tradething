package service

import (
	"context"
	"errors"
	"log"
	"time"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"

	utils "github.com/non26/tradepkg/pkg/bn/utils"
)

// NOTE:
//1. this function is used to place a single order for buying long/ selling short to open position
// right now it only support for stop loss and take profit wwatch with "hour timeframe" in intraday .
//2. close position is not support partially closed
//3. for accumulate order, client are resposible for calculating the amount of qty to be placed in each bar
//4. for watching order, it'll be triggered when the bar closed

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *svcFuture.PlaceSignleOrderServiceRequest,
) (*bnfutureres.PlaceSignleOrderHandlerResponse, error) {

	positionHistory, err := bfs.repository.GetHistoryByClientID(ctx, request.GetClientOrderId())
	if err != nil {
		log.Println("error get history by client id", err.Error())
		return nil, err
	}
	if positionHistory.IsFound() {
		return nil, errors.New("position client id is not valid")
	}

	openingPositionTable := request.ToBinanceFutureOpeningPositionRepositoryModel()
	dbOpeningOrder, err := bfs.repository.GetOpenOrderByKey(ctx, openingPositionTable.GetKeyByPositionSideAndSymbol())
	if err != nil {
		log.Println("error get open order by key", err.Error())
		return nil, err
	}

	if utils.IsSellCrypto(request.GetSide(), request.GetPositionSide()) {
		placeSellOrderRes, err := bfs.binanceService.PlaceSingleOrder(
			ctx,
			request.ToBinanceServiceModel(),
		)
		if err != nil {
			log.Println("error place sell order", err.Error())
			return nil, err
		}
		bfs.repository.DeleteOpenOrderByKey(ctx, openingPositionTable.GetKeyByPositionSideAndSymbol())
		bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
		return placeSellOrderRes.ToBnHandlerResponse(), nil
	}

	if request.IsStopLossNil() {
		return nil, errors.New("stop loss is mandatory")
	}

	if !dbOpeningOrder.IsFound() { // meaing this is new order, no existing order is found
		dbQUsdt, err := bfs.repository.GetQouteUSDT(ctx, request.GetSymbol())
		if err != nil {
			log.Println("error get qoute usdt", err.Error())
			return nil, err
		}

		if !dbQUsdt.IsFound() {
			dbQUsdt.SetSymbol(request.GetSymbol())
			dbQUsdt.SetCountingLong(0)
			dbQUsdt.SetCountingShort(0)
			err = bfs.repository.InsertNewSymbolQouteUSDT(ctx, dbQUsdt)
			if err != nil {
				log.Println("error insert new symbol qoute usdt", err.Error())
				return nil, err
			}
		}

		if request.GetClientOrderId() == "" {
			var counting int
			if request.GetPositionSide() == bfs.positionSideType.Short() {
				counting = dbQUsdt.GetCountingShort()
			} else {
				counting = dbQUsdt.GetCountingLong()
			}
			request.SetClientOrderId(utils.BinanceDefaultClientID(request.GetSymbol(), request.GetPositionSide(), counting))
		}

		placeOrderRes, err := bfs.binanceService.PlaceSingleOrder(
			ctx,
			request.ToBinanceServiceModel(),
		)
		if err != nil {
			log.Println("error place order", err.Error())
			return nil, err
		}

		if request.GetPositionSide() == bfs.positionSideType.Short() {
			dbQUsdt.SetCountingShort(dbQUsdt.GetCountingShort() + 1)
		} else {
			dbQUsdt.SetCountingLong(dbQUsdt.GetCountingLong() + 1)
		}
		err = bfs.repository.UpdateQouteUSDT(ctx, dbQUsdt)
		if err != nil {
			log.Println("error update qoute usdt", err.Error())
		}
		err = bfs.repository.NewOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())
		if err != nil {
			log.Println("error new open order", err.Error())
		}

		return placeOrderRes.ToBnHandlerResponse(), nil
	} else { // there is existing order
		if request.GetSymbol() != dbOpeningOrder.Symbol {
			return nil, errors.New("symbol not match")
		}
		if request.GetPositionSide() != dbOpeningOrder.PositionSide {
			return nil, errors.New("position side not match")
		}
		if request.GetClientOrderId() == "" {
			return nil, errors.New("client order id is empty")
		}

		if request.GetClientOrderId() != dbOpeningOrder.ClientId {
			// for accumulate order
			placeOrderRes, err := bfs.binanceService.PlaceSingleOrder(
				ctx,
				request.ToBinanceServiceModel(),
			)
			if err != nil {
				log.Println("error place order for accumulate order", err.Error())
				return nil, err
			}
			dbOpeningOrder.AddMoreAmountQ(request.GetAmountQ())
			err = bfs.repository.UpdateOpenOrder(ctx, dbOpeningOrder)
			if err != nil {
				log.Println("error update open order for accumulate order", err.Error())
				return nil, err
			}
			return placeOrderRes.ToBnHandlerResponse(), nil
		} else {
			// for watching order
			bnTime := utils.NewBinanceTime(time.Now())

			period, unit, err := utils.GetInterval(request.GetStopLoss().Interval)
			if err != nil {
				log.Println("error get interval for watching order", err.Error())
				return nil, err
			}
			var prv_start, prv_end time.Time
			/// now support only hourly
			switch unit {
			case utils.Minute:
				var err error
				prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartMinuteAndEndMinute(period)
				if err != nil {
					log.Println("error get previous bn time start minute and end minute for watching order", err.Error())
					return nil, err
				}
			case utils.Hour:
				var err error
				prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartHourAndEndHour(period)
				if err != nil {
					log.Println("error get previous bn time start hour and end hour for watching order", err.Error())
					return nil, err
				}
			case utils.Day:
				var err error
				prv_start, prv_end, err = bnTime.GetPreviousBnTimeStartDayAndEndDay(period)
				if err != nil {
					log.Println("error get previous bn time start day and end day for watching order", err.Error())
					return nil, err
				}
			default:
				return nil, errors.New("invalid interval")
			}

			dbMarketData, err := bfs.bnMarketDataService.GetCandleStickData(ctx, request.ToBnCandleStickModel(
				utils.GetSpecificBnTimestamp(&prv_start),
				utils.GetSpecificBnTimestamp(&prv_end),
			))
			if err != nil {
				log.Println("error get candle stick data for watching order", err.Error())
				return nil, err
			}

			if !request.IsStopLossNil() {
				if request.GetPositionSide() == bfs.positionSideType.Long() {
					if dbMarketData.GetClosePrice().GetFloat64() < request.GetStopLoss().Price {
						request.SetSide(bfs.sideType.Sell())
						placeSellOrderRes, err := bfs.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
						if err != nil {
							log.Println("error place sell order for watching order", err.Error())
							return nil, err
						}
						bfs.repository.DeleteOpenOrderByKey(ctx, dbOpeningOrder.GetKeyByPositionSideAndSymbol())
						bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
						return placeSellOrderRes.ToBnHandlerResponse(), nil
					}
				} else if request.GetPositionSide() == bfs.positionSideType.Short() {
					if dbMarketData.GetClosePrice().GetFloat64() > request.GetStopLoss().Price {
						request.SetSide(bfs.sideType.Buy())
						placeBuyOrderRes, err := bfs.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
						if err != nil {
							log.Println("error place buy order for watching order", err.Error())
							return nil, err
						}
						bfs.repository.DeleteOpenOrderByKey(ctx, dbOpeningOrder.GetKeyByPositionSideAndSymbol())
						bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
						return placeBuyOrderRes.ToBnHandlerResponse(), nil
					}
				}
			}

			if !request.IsTakeProfitNil() {
				if request.GetPositionSide() == bfs.positionSideType.Long() {
					if dbMarketData.GetClosePrice().GetFloat64() > request.GetTakeProfit().Price {
						request.SetSide(bfs.sideType.Sell())
						placeSellOrderRes, err := bfs.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
						if err != nil {
							log.Println("error place sell order for watching order", err.Error())
							return nil, err
						}
						bfs.repository.DeleteOpenOrderByKey(ctx, dbOpeningOrder.GetKeyByPositionSideAndSymbol())
						bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
						return placeSellOrderRes.ToBnHandlerResponse(), nil
					}
				} else if request.GetPositionSide() == bfs.positionSideType.Short() {
					if dbMarketData.GetClosePrice().GetFloat64() < request.GetTakeProfit().Price {
						request.SetSide(bfs.sideType.Buy())
						placeBuyOrderRes, err := bfs.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
						if err != nil {
							log.Println("error place buy order for watching order", err.Error())
							return nil, err
						}
						bfs.repository.DeleteOpenOrderByKey(ctx, dbOpeningOrder.GetKeyByPositionSideAndSymbol())
						bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
						return placeBuyOrderRes.ToBnHandlerResponse(), nil
					}
				}
			}

		}
	}

	return nil, nil
}

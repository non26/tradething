package service

import (
	"context"
	"errors"
	"time"
	bnfutureres "tradething/app/bn/bn_future/handler_response_model"
	svcFuture "tradething/app/bn/bn_future/service_model"

	"github.com/non26/tradepkg/pkg/bn/thaitime"
	utils "github.com/non26/tradepkg/pkg/bn/utils"
)

func (bfs *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *svcFuture.PlaceSignleOrderServiceRequest,
) (*bnfutureres.PlaceSignleOrderHandlerResponse, error) {

	positionHistory, err := bfs.repository.GetHistoryByClientID(ctx, request.GetClientOrderId())
	if err != nil {
		return nil, err
	}
	if positionHistory.IsFound() {
		return nil, errors.New("position history already exist")
	}

	openingPositionTable := request.ToBinanceFutureOpeningPositionRepositoryModel()
	dbOpeningOrder, err := bfs.repository.GetOpenOrderByKey(ctx, openingPositionTable.GetKeyByPositionSideAndSymbol())
	if err != nil {
		return nil, err
	}

	if utils.IsSellCrypto(request.GetSide(), request.GetPositionSide()) {
		placeSellOrderRes, err := bfs.binanceService.PlaceSingleOrder(
			ctx,
			request.ToBinanceServiceModel(),
		)
		if err != nil {
			return nil, err
		}
		bfs.repository.DeleteOpenOrderByKey(ctx, openingPositionTable.GetKeyByPositionSideAndSymbol())
		bfs.repository.InsertHistory(ctx, request.ToBnPositionHistoryRepositoryModel())
		return placeSellOrderRes.ToBnHandlerResponse(), nil
	}

	if request.IsStopLossNil() {
		return nil, errors.New("stop loss is nil")
	}

	if !dbOpeningOrder.IsFound() { // meaing this is new order, no existing order is found
		dbQUsdt, err := bfs.repository.GetQouteUSDT(ctx, request.GetSymbol())
		if err != nil {
			return nil, err
		}

		if !dbQUsdt.IsFound() {
			dbQUsdt.SetSymbol(request.GetSymbol())
			dbQUsdt.SetCountingLong(0)
			dbQUsdt.SetCountingShort(0)
			err = bfs.repository.InsertNewSymbolQouteUSDT(ctx, dbQUsdt)
			if err != nil {
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
			return nil, err
		}

		if request.GetPositionSide() == bfs.positionSideType.Short() {
			dbQUsdt.SetCountingShort(dbQUsdt.GetCountingShort() + 1)
		} else {
			dbQUsdt.SetCountingLong(dbQUsdt.GetCountingLong() + 1)
		}
		bfs.repository.UpdateQouteUSDT(ctx, dbQUsdt)
		bfs.repository.NewOpenOrder(ctx, request.ToBinanceFutureOpeningPositionRepositoryModel())

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
				return nil, err
			}
			dbOpeningOrder.AddMoreAmountQ(request.GetAmountQ())
			bfs.repository.UpdateOpenOrder(ctx, dbOpeningOrder)
			return placeOrderRes.ToBnHandlerResponse(), nil
		} else {
			// for watching order
			thaitime, err := thaitime.NewThaiTime()
			if err != nil {
				return nil, err
			}

			period, unit, err := utils.GetInterval(request.GetStopLoss().Interval)
			if err != nil {
				return nil, err
			}

			start, end := utils.GetPreviousUnixBnStartAndEndOfPeriodHours(time.Now(), thaitime, period, unit)
			dbMarketData, err := bfs.bnMarketDataService.GetCandleStickData(ctx, request.ToBnCandleStickModel(
				utils.GetSpecificBnTimestamp(start),
				utils.GetSpecificBnTimestamp(end),
			))
			if err != nil {
				return nil, err
			}

			if !request.IsStopLossNil() {
				if request.GetPositionSide() == bfs.positionSideType.Long() {
					if dbMarketData.GetClosePrice().GetFloat64() < request.GetStopLoss().Price {
						request.SetSide(bfs.sideType.Sell())
						placeSellOrderRes, err := bfs.binanceService.PlaceSingleOrder(ctx, request.ToBinanceServiceModel())
						if err != nil {
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

package service

import (
	"context"
	"errors"
	handlerres "tradething/app/bn/bn_future/handler_response"
	model "tradething/app/bn/bn_future/service_model"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	serviceerror "github.com/non26/tradepkg/pkg/bn/service_error"

	utils "github.com/non26/tradepkg/pkg/bn/utils"
)

// NOTE:
//1. this function is used to place a single order for buying long/ selling short to open position
// right now it only support for stop loss and take profit wwatch with "hour timeframe" in intraday .
//2. close position is not support partially closed
//3. for accumulate order, client are resposible for calculating the amount of qty to be placed in each bar
//4. for watching order, it'll be triggered when the bar closed

func (b *binanceFutureService) PlaceSingleOrder(
	ctx context.Context,
	request *model.Position,
) (*handlerres.PlacePosition, serviceerror.IError) {

	positionHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientId())
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, err)
	}
	if positionHistory.IsFound() {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_HISTORY_ERROR, errors.New("position client id is not valid"))
	}

	openingPositionTable := request.ToBinanceFutureOpeningPositionRepositoryModel()
	dbOpeningPosition, err := b.bnFtOpeningPositionTable.Get(ctx, openingPositionTable)
	if err != nil {
		return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
	}

	if !(dbOpeningPosition.IsFound() || utils.IsSellPosition(request.GetPositionSide(), request.GetSide())) {
		// meaing this is new order, no existing order is found
		placeOrderRes, svcerr := b.openPosition(ctx, request)
		if svcerr != nil {
			return nil, svcerr
		}
		return placeOrderRes.ToBnHandlerResponse(), nil
	} else {
		// there is existing order
		if request.IsSellOrder() {
			placeSellOrderRes, svcerr := b.closePosition(ctx, request)
			if svcerr != nil {
				return nil, svcerr
			}
			return placeSellOrderRes.ToBnHandlerResponse(), nil
		}

		if request.GetSymbol() != dbOpeningPosition.Symbol {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("symbol not match"))
		}
		if request.GetPositionSide() != dbOpeningPosition.PositionSide {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("position side not match"))
		}
		if request.GetClientId() == "" {
			return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, errors.New("client order id is empty"))
		}

		if request.GetClientId() != dbOpeningPosition.ClientId {
			placeOrderRes, svcerr := b.accumulateOrder(ctx, request, dbOpeningPosition)
			if svcerr != nil {
				return nil, svcerr
			}
			return placeOrderRes.ToBnHandlerResponse(), nil
		} else {
			prv_start, prv_end, svcerr := b.getPreviousBnTimeStartAndEnd(request)
			if svcerr != nil {
				return nil, svcerr
			}

			dbMarketData, err := b.bnMarketDataService.GetCandleStickData(ctx, request.ToBnCandleStickModel(
				utils.GetSpecificBnTimestamp(prv_start),
				utils.GetSpecificBnTimestamp(prv_end),
			))
			if err != nil {
				return nil, serviceerror.NewServiceErrorWith(serviceerror.BN_OPENING_POSITION_ERROR, err)
			}

			closePrice := dbMarketData.GetClosePrice().GetFloat64()
			if !request.IsStopLossNil() {
				if request.IsLongPosition() {
					if closePrice < request.GetStopLoss().Price {
						request.SetSide(bnconstant.SELL)
						closePositionRes, svcerr := b.closePosition(ctx, request)
						if svcerr != nil {
							return nil, svcerr
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				} else if request.IsShortPosition() {
					if closePrice > request.GetStopLoss().Price {
						request.SetSide(bnconstant.BUY)
						closePositionRes, svcerr := b.closePosition(ctx, request)
						if svcerr != nil {
							return nil, svcerr
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				}
			}

			if !request.IsTakeProfitNil() {
				if request.IsLongPosition() {
					if dbMarketData.GetClosePrice().GetFloat64() > request.GetTakeProfit().Price {
						request.SetSide(bnconstant.SELL)
						closePositionRes, svcerr := b.closePosition(ctx, request)
						if svcerr != nil {
							return nil, svcerr
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				} else if request.IsShortPosition() {
					if dbMarketData.GetClosePrice().GetFloat64() < request.GetTakeProfit().Price {
						request.SetSide(bnconstant.BUY)
						closePositionRes, svcerr := b.closePosition(ctx, request)
						if svcerr != nil {
							return nil, svcerr
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				}
			}

		}
	}

	return nil, nil
}

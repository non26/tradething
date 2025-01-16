package service

import (
	"context"
	"errors"
	"log"
	handlerres "tradething/app/bn/bn_future/handler_response_model"
	model "tradething/app/bn/bn_future/service_model"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"

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
) (*handlerres.PlacePosition, error) {

	positionHistory, err := b.bnFtHistoryTable.Get(ctx, request.GetClientOrderId())
	if err != nil {
		log.Println("error get history by client id", err.Error())
		return nil, err
	}
	if positionHistory.IsFound() {
		return nil, errors.New("position client id is not valid")
	}

	openingPositionTable := request.ToBinanceFutureOpeningPositionRepositoryModel()
	dbOpeningOrder, err := b.bnFtOpeningPositionTable.Get(ctx, openingPositionTable)
	if err != nil {
		log.Println("error get open order by key", err.Error())
		return nil, err
	}

	if request.IsSellOrder() {
		placeSellOrderRes, err := b.closePosition(ctx, request)
		if err != nil {
			log.Println("error close position", err.Error())
			return nil, err
		}
		return placeSellOrderRes.ToBnHandlerResponse(), nil
	}

	if request.IsStopLossNil() {
		return nil, errors.New("stop loss is mandatory")
	}

	if !dbOpeningOrder.IsFound() { // meaing this is new order, no existing order is found
		dbQUsdt, err := b.bnFtQouteUsdtTable.Get(ctx, request.GetSymbol())
		if err != nil {
			log.Println("error get qoute usdt", err.Error())
			return nil, err
		}

		if !dbQUsdt.IsFound() {
			dbQUsdt = dynamodbmodel.NewBinanceFutureQouteUSTDTableRecord(request.GetSymbol(), request.IsLongPosition())
			err = b.bnFtQouteUsdtTable.Insert(ctx, dbQUsdt)
			if err != nil {
				log.Println("error insert new symbol qoute usdt", err.Error())
				return nil, err
			}
		}

		// Set Default Client Order Id
		if request.GetClientOrderId() == "" {
			b.setDefaultClientOrderId(request, dbQUsdt)
		}

		placeOrderRes, err := b.openPosition(ctx, request, dbQUsdt)
		if err != nil {
			log.Println("error open position", err.Error())
			return nil, err
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
			placeOrderRes, err := b.accumulateOrder(ctx, request, dbOpeningOrder)
			if err != nil {
				log.Println("error accumulate order", err.Error())
				return nil, err
			}
			return placeOrderRes.ToBnHandlerResponse(), nil
		} else {
			prv_start, prv_end, err := b.getPreviousBnTimeStartAndEnd(request)
			if err != nil {
				log.Println("error get previous bn time start and end for watching order", err.Error())
				return nil, err
			}

			dbMarketData, err := b.bnMarketDataService.GetCandleStickData(ctx, request.ToBnCandleStickModel(
				utils.GetSpecificBnTimestamp(prv_start),
				utils.GetSpecificBnTimestamp(prv_end),
			))
			if err != nil {
				log.Println("error get candle stick data for watching order", err.Error())
				return nil, err
			}
			closePrice := dbMarketData.GetClosePrice().GetFloat64()
			if !request.IsStopLossNil() {
				if request.IsLongPosition() {
					if closePrice < request.GetStopLoss().Price {
						request.SetSide(b.sideType.Sell())
						closePositionRes, err := b.closePosition(ctx, request)
						if err != nil {
							log.Println("error close position", err.Error())
							return nil, err
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				} else if request.IsShortPosition() {
					if closePrice > request.GetStopLoss().Price {
						request.SetSide(b.sideType.Buy())
						closePositionRes, err := b.closePosition(ctx, request)
						if err != nil {
							log.Println("error close position", err.Error())
							return nil, err
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				}
			}

			if !request.IsTakeProfitNil() {
				if request.IsLongPosition() {
					if dbMarketData.GetClosePrice().GetFloat64() > request.GetTakeProfit().Price {
						request.SetSide(b.sideType.Sell())
						closePositionRes, err := b.closePosition(ctx, request)
						if err != nil {
							log.Println("error close position", err.Error())
							return nil, err
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				} else if request.IsShortPosition() {
					if dbMarketData.GetClosePrice().GetFloat64() < request.GetTakeProfit().Price {
						request.SetSide(b.sideType.Buy())
						closePositionRes, err := b.closePosition(ctx, request)
						if err != nil {
							log.Println("error close position", err.Error())
							return nil, err
						}
						return closePositionRes.ToBnHandlerResponse(), nil
					}
				}
			}

		}
	}

	return nil, nil
}

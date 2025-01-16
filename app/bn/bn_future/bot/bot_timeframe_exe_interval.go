package bnfuture

import (
	"context"
	"errors"
	"log"
	"time"
	handlerres "tradething/app/bn/bn_future/bot_handler_response"
	bnsvcreq "tradething/app/bn/bn_future/bot_model"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func isBetweenTime(startDate, endDate time.Time, presentTime time.Time) bool {
	if startDate.Unix() <= presentTime.Unix() && presentTime.Unix() <= endDate.Unix() {
		return true
	}
	return false
}

func (b *botService) BotTimeframeExeInterval(ctx context.Context, req *bnsvcreq.BotTimeframeExeIntervalRequest) (*handlerres.BotTimeframeExeIntervalResponse, error) {
	presentTime := time.Now().UTC()
	inTime := isBetweenTime(req.GetStartDate(), req.GetEndDate(), presentTime)

	bot, err := b.bnFtBotTable.Get(ctx, req.GetBotId())
	if err != nil {
		return nil, err
	}

	if !bot.IsFound() {
		return nil, errors.New("bot not found")
	}

	posHistory, err := b.bnFtHistoryTable.Get(ctx, req.GetBotOrderID())
	if err != nil {
		return nil, err
	}

	if posHistory.IsFound() {
		return nil, errors.New("bot order already closed")
	}

	current_position, err := b.bnFtBotOnRunTable.Get(ctx, req.ToBnFtBotOnRun())
	if err != nil {
		return nil, errors.New("get current position error")
	}

	var closeSide string
	if current_position.IsFound() {
		if req.GetBotOrderID() != current_position.BotOrderID {
			return nil, errors.New("bot order already exists")
		}

		if req.GetPositionSide() != current_position.PositionSide {
			return nil, errors.New("position side not match")
		}
		if current_position.PositionSide == b.positionSideType.Long() {
			closeSide = b.sideType.Sell()
		} else {
			closeSide = b.sideType.Buy()
		}
	}

	var openSide string
	if req.GetPositionSide() == b.positionSideType.Long() {
		openSide = b.sideType.Buy()
	} else {
		openSide = b.sideType.Sell()
	}

	if inTime {
		if current_position.IsFound() {
			if !current_position.IsActive {
				return nil, errors.New("bot order already not active")
			}
			// close current position
			closeOrder := req.ToBnFtPlaceSingleOrderServiceRequest(closeSide, b.orderType.Market())
			closeOrder.EntryQuantity = current_position.AmountQoute
			_, err := b.binanceService.PlaceSingleOrder(ctx, closeOrder)
			if err != nil {
				return nil, errors.New("place order error")
			}
			// err = b.repository.DeleteBotOnRun(ctx, req.ToBnFtDeleteBotOnRun())
			// if err != nil {
			// 	log.Println("delete bot on run error", err)
			// }
		}
		// open new position
		_, err = b.binanceService.PlaceSingleOrder(ctx, req.ToBnFtPlaceSingleOrderServiceRequest(openSide, b.orderType.Market()))
		if err != nil {
			log.Println("place order error", err)
			return nil, errors.New("place order error")
		}

		isFirstTime := !current_position.IsFound()
		if isFirstTime {
			err = b.bnFtBotOnRunTable.Insert(ctx, req.ToBnFtBotOnRun())
			if err != nil {
				log.Println("insert bot on run error", err)
			}
		} else {
			err = b.bnFtBotOnRunTable.Update(ctx, req.ToBnFtBotOnRun())
			if err != nil {
				log.Println("update bot on run error", err)
			}
		}

		qouteUSDT, err := b.bnFtQouteUsdtTable.Get(ctx, req.GetSymbol())
		if err != nil {
			log.Println("get qoute usdt error", err)
		}

		if !qouteUSDT.IsFound() {
			qouteUSDT = dynamodbmodel.NewBnFtQouteUSDT()
			qouteUSDT.Symbol = req.GetSymbol()
			if req.GetPositionSide() == b.positionSideType.Long() {
				qouteUSDT.SetCountingLong(qouteUSDT.GetNextCountingLong().Int())
			} else {
				qouteUSDT.SetCountingShort(qouteUSDT.GetNextCountingShort().Int())
			}
			err = b.bnFtQouteUsdtTable.Insert(ctx, qouteUSDT)
			if err != nil {
				log.Println("insert qoute usdt error", err)
			}
		} else {
			qouteUSDT.Symbol = req.GetSymbol()
			if req.GetPositionSide() == b.positionSideType.Long() {
				qouteUSDT.SetCountingLong(qouteUSDT.GetNextCountingLong().Int())
			} else {
				qouteUSDT.SetCountingShort(qouteUSDT.GetNextCountingShort().Int())
			}
			err = b.bnFtQouteUsdtTable.Update(ctx, qouteUSDT)
			if err != nil {
				log.Println("update qoute usdt error", err)
			}
		}

	} else { // off time
		if current_position.IsFound() {
			// close position
			closeOrder := req.ToBnFtPlaceSingleOrderServiceRequest(closeSide, b.orderType.Market())
			closeOrder.EntryQuantity = current_position.AmountQoute
			_, err := b.binanceService.PlaceSingleOrder(ctx, closeOrder)
			if err != nil {
				return nil, errors.New("place order error")
			}
			err = b.bnFtBotOnRunTable.Delete(ctx, req.ToBnFtDeleteBotOnRun())
			if err != nil {
				log.Println("delete bot on run error", err)
			}
			err = b.bnFtHistoryTable.Insert(ctx, req.ToBnFtHistory())
			if err != nil {
				log.Println("insert history error", err)
			}
		}
		return &handlerres.BotTimeframeExeIntervalResponse{
			BotOrderID: req.GetBotOrderID(),
			Status:     "success",
			Message:    "no bot open",
		}, nil
	}

	return &handlerres.BotTimeframeExeIntervalResponse{
		BotOrderID: req.GetBotOrderID(),
		Status:     "success",
		Message:    "success",
	}, nil
}

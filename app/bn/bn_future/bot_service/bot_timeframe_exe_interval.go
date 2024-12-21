package bnfuture

import (
	"log"
	handlerres "tradething/app/bn/bn_future/bot_handler_response_model"
	bnsvcreq "tradething/app/bn/bn_future/bot_service_model"

	"context"
	"errors"
	"time"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_repository/models"
)

func isBetweenTime(startDate, endDate time.Time, presentTime time.Time) bool {
	if startDate.Unix() <= presentTime.Unix() && endDate.Unix() >= presentTime.Unix() {
		return true
	}
	return false
}

func (b *botService) BotTimeframeExeInterval(ctx context.Context, req *bnsvcreq.BotTimeframeExeIntervalRequest) (*handlerres.BotTimeframeExeIntervalResponse, error) {
	presentTime := time.Now().UTC()
	inTime := isBetweenTime(req.GetStartDate(), req.GetEndDate(), presentTime)

	posHistory, err := b.repository.GetHistoryByClientID(context.Background(), req.GetBotOrderID())
	if err != nil {
		return nil, err
	}

	if posHistory.IsFound() {
		return nil, errors.New("bot order already exists")
	}

	current_position := req.ToBnFtOpeningPosition()
	current_position, err = b.repository.GetOpenOrderBySymbolAndPositionSide(ctx, current_position)
	if err != nil {
		return nil, errors.New("get current position error")
	}

	var closeSide string
	if current_position.IsFound() {
		if req.GetBotOrderID() != current_position.ClientId {
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
			// close current position
			_, err := b.binanceService.PlaceSingleOrder(ctx, req.ToBnFtPlaceSingleOrderServiceRequest(closeSide, b.orderType.Market()))
			if err != nil {
				return nil, errors.New("place order error")
			}
			err = b.repository.DeleteOpenOrderBySymbolAndPositionSide(ctx, current_position)
			if err != nil {
				log.Println("delete current position error", err)
			}
			err = b.repository.DeleteBotOnRun(ctx, req.ToBnFtDeleteBotOnRun())
			if err != nil {
				log.Println("delete bot on run error", err)
			}
			err = b.repository.InsertHistory(ctx, req.ToBnFtHistory())
			if err != nil {
				log.Println("insert history error", err)
			}
		}
		// open new position
		_, err = b.binanceService.PlaceSingleOrder(ctx, req.ToBnFtPlaceSingleOrderServiceRequest(openSide, b.orderType.Market()))
		if err != nil {
			log.Println("place order error", err)
		}
		err = b.repository.InsertNewOpenOrder(ctx, req.ToBnFtOpeningPosition())
		if err != nil {
			log.Println("insert open position error", err)
		}
		err = b.repository.InsertBotOnRun(ctx, req.ToBnFtBotOnRun())
		if err != nil {
			log.Println("insert bot on run error", err)
		}
		qouteUSDT := dynamodbmodel.NewBnFtQouteUSDT()
		qouteUSDT.Symbol = req.GetSymbol()
		if current_position.PositionSide == req.GetPositionSide() {
			if current_position.PositionSide == b.positionSideType.Long() {
				qouteUSDT.GetNextCountingLong()
			} else {
				qouteUSDT.GetNextCountingShort()
			}
		}
		err = b.repository.UpdateQouteUSDT(ctx, qouteUSDT)
		if err != nil {
			log.Println("update qoute usdt error", err)
		}
	} else { // off time
		if current_position.IsFound() {
			// close position
			_, err := b.binanceService.PlaceSingleOrder(ctx, req.ToBnFtPlaceSingleOrderServiceRequest(closeSide, b.orderType.Market()))
			if err != nil {
				return nil, errors.New("place order error")
			}
			err = b.repository.DeleteOpenOrderBySymbolAndPositionSide(ctx, current_position)
			if err != nil {
				log.Println("delete current position error", err)
			}
			err = b.repository.DeleteBotOnRun(ctx, req.ToBnFtDeleteBotOnRun())
			if err != nil {
				log.Println("delete bot on run error", err)
			}
			err = b.repository.InsertHistory(ctx, req.ToBnFtHistory())
			if err != nil {
				log.Println("insert history error", err)
			}
		}
	}

	return &handlerres.BotTimeframeExeIntervalResponse{
		BotOrderID: req.GetBotOrderID(),
		Status:     "success",
		Message:    "success",
	}, nil
}

package bnfuture

// import (
// 	"context"
// 	bnservice "tradething/app/bn/bn_future/bnservice"
// 	bothandlerreq "tradething/app/bn/bn_future/bot_handler_request"
// 	bothandlerres "tradething/app/bn/bn_future/bot_handler_response"

// 	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_repository"
// 	positionconstant "github.com/non26/tradepkg/pkg/bn/position_constant"
// )

// type IBotService interface {
// 	TimeIntervalSemiBotService(
// 		ctx context.Context,
// 		req *bothandlerreq.TradeTimeIntervalBinanceFutureRequest,
// 	) (*bothandlerres.TradeTimeIntervalBinanceFutureResponse, error)
// }

// type botService struct {
// 	bn_service    bnservice.IBinanceFutureExternalService
// 	bn_repository bndynamodb.IRepository
// 	order_type    positionconstant.IOrderType
// 	position_side positionconstant.IPositionSide
// 	side          positionconstant.ISide
// }

// func NewBotService(
// 	bn_service bnservice.IBinanceFutureExternalService,
// 	bn_repository bndynamodb.IRepository,
// 	order_type positionconstant.IOrderType,
// 	position_side positionconstant.IPositionSide,
// 	side positionconstant.ISide,
// ) IBotService {
// 	return &botService{
// 		bn_service:    bn_service,
// 		bn_repository: bn_repository,
// 		order_type:    order_type,
// 		position_side: position_side,
// 		side:          side,
// 	}
// }

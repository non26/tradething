package req

import (
	"tradething/app/bn/future/BFF/trade/domain"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type NewOrderReq struct {
	AccountId        string `json:"accountId" binding:"required"`
	PositionSide     string `json:"positionSide" binding:"required"`
	Side             string `json:"side" binding:"required"`
	Quantity         string `json:"quantity" binding:"required"`
	Symbol           string `json:"symbol" binding:"required"`
	NewClientOrderId string `json:"newClientOrderId" binding:"required"`
	Type             string `json:"type" binding:"required"`
}

func (r *NewOrderReq) ToTradeAdaptorNewOrderRequest(d *domain.Order) *NewOrderReq {
	return &NewOrderReq{
		AccountId:        d.AccountId,
		PositionSide:     d.PositionSide,
		Side:             d.Side,
		Quantity:         d.AmountB,
		Symbol:           d.Symbol,
		NewClientOrderId: d.ClientId,
		Type:             bnconstant.MARKET,
	}
}

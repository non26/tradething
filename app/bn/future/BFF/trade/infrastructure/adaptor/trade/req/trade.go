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

func (r *NewOrderReq) ToTradeAdaptorNewOrderRequest(d *domain.Order) {
	r.AccountId = d.AccountId
	r.PositionSide = d.PositionSide
	r.Side = d.Side
	r.Quantity = d.AmountB
	r.Symbol = d.Symbol
	r.NewClientOrderId = d.ClientId
	r.Type = bnconstant.MARKET

}

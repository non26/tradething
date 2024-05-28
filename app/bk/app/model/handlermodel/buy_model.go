package model

import model "tradetoolv2/app/bk/app/model/bitkubservicemodel"

type BuyOrderHandlerRequest struct {
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount"`
	Ratio    float64 `json:"ratio"`
	Type     string  `json:"type"`
	ClientId string  `json:"client_id"`
}

func (b *BuyOrderHandlerRequest) ToBuyOrderBkServiceRequest() *model.BuyOrderBkServiceRequest {
	d := model.BuyOrderBkServiceRequest{
		Symbol:   b.Symbol,
		Amount:   b.Amount,
		Ratio:    b.Ratio,
		Type:     b.Type,
		ClientId: b.ClientId,
	}
	return &d
}

type BuyOrderHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

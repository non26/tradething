package model

import model "tradething/app/bk/app/model/bitkubservicemodel"

type SellHandlerRequest struct {
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount"`
	Ratio    float64 `json:"ratio"`
	Type     string  `json:"type"`
	ClientId string  `json:"client_id"`
}

func (s *SellHandlerRequest) ToSellOrderBkRequest() *model.SellBkServiceRequest {
	d := model.SellBkServiceRequest{
		Symbol:   s.Symbol,
		Amount:   s.Amount,
		Ratio:    s.Ratio,
		Type:     s.Type,
		ClientId: s.ClientId,
	}
	return &d
}

type SellOrderHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

package model

type BuyOrderBkServiceRequest struct {
	Symbol   string  `json:"sym"`
	Amount   float64 `json:"amt"`
	Ratio    float64 `json:"rat"`
	Type     string  `json:"typ"`
	ClientId string  `json:"client_id"`
}

type BuyOrderBkServiceResponse struct {
	Id            string  `json:"id"`   // order id
	Hash          string  `json:"hash"` // order hash
	Type          string  `json:"typ"`  // order type
	Amount        float64 `json:"amt"`  // spending amount
	Ratio         float64 `json:"rat"`  // rate
	Fee           float64 `json:"fee"`  // fee
	Credit        float64 `json:"cre"`  // fee credit used
	AmountReceive float64 `json:"rec"`  // amount to receive
	Timestamp     float64 `json:"ts"`   // timestamp
	RefId         string  `json:"ci"`   // input id for reference
}

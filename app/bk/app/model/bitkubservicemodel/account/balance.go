package model

type BalanceBkServiceDataResponse struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type BalanceBkServiceResponse struct {
	Symbol map[string]BalanceBkServiceDataResponse
}

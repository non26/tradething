package req

import "errors"

type SetAdvancedPositionReqs struct {
	Positions []SetAdvancedPositionReq `json:"positions"`
}

type SetAdvancedPositionReq struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	PositionSide string `json:"position_side"`
	AmountB      string `json:"amount_b"`
}

func (r *SetAdvancedPositionReq) Validate() error {
	if r.ClientId == "" {
		return errors.New("client_id is required")
	}
	if r.Symbol == "" {
		return errors.New("symbol is required")
	}
	if r.Side == "" {
		return errors.New("side is required")
	}
	if r.PositionSide == "" {
		return errors.New("position_side is required")
	}
	if r.AmountB == "" {
		return errors.New("amount_b is required")
	}
	return nil
}

package res

import "tradething/app/bn/future/position_history/domain"

type GetHistoryRes struct {
	ClientId     string `json:"client_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
	AccountId    string `json:"account_id"`
}

func (r *GetHistoryRes) FromDomain(history *domain.BnFtHistory) *GetHistoryRes {
	if history == nil {
		return nil
	}
	return &GetHistoryRes{
		ClientId:     history.ClientId,
		Symbol:       history.Symbol,
		PositionSide: history.PositionSide,
		AccountId:    history.AccountId,
	}
}

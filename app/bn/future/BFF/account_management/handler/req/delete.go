package req

import "tradething/app/bn/future/BFF/account_management/domain"

type DeleteReq struct {
	AccountId string `json:"accountId"`
}

func (r *DeleteReq) ToDomain() *domain.Account {
	return &domain.Account{
		AccountId: r.AccountId,
	}
}

package req

import "tradething/app/bn/future/sub_account/domain"

type UpsertReq struct {
	AccountId   string `json:"accountId"`
	AccountName string `json:"accountName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

func (r *UpsertReq) ToDomain() *domain.SubAccount {
	return &domain.SubAccount{
		AccountId:   r.AccountId,
		AccountName: r.AccountName,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
	}
}

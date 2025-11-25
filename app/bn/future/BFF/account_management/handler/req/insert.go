package req

import "tradething/app/bn/future/BFF/account_management/domain"

type InsertReq struct {
	AccountId   string `json:"accountId"`
	AccountName string `json:"accountName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

func (r *InsertReq) ToDomain() *domain.Account {
	return &domain.Account{
		AccountId:        r.AccountId,
		AccountName:      r.AccountName,
		AccountStartDate: r.StartDate,
		AccountEndDate:   r.EndDate,
	}
}

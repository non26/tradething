package domain

import (
	"tradething/app/bn/future/sub_account/domain"
)

type Account struct {
	AccountId        string
	AccountName      string
	AccountStartDate string
	AccountEndDate   string
}

func (a *Account) ToSubAccountExtSrv() *domain.SubAccount {
	return &domain.SubAccount{
		AccountId:   a.AccountId,
		AccountName: a.AccountName,
		StartDate:   a.AccountStartDate,
		EndDate:     a.AccountEndDate,
	}
}

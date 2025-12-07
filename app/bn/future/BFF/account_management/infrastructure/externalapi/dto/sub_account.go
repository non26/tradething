package dto

import (
	"tradething/app/bn/future/BFF/account_management/domain"

	subAccountDomain "tradething/app/bn/future/sub_account/domain"
)

// import "tradething/app/bn/future/sub_account/domain"

type SubAccountDto struct {
	AccountId   string
	AccountName string
	StartDate   string
	EndDate     string
	IsExpired   bool
}

func NewSubAccountDto() *SubAccountDto {
	return &SubAccountDto{}
}

func (s *SubAccountDto) FromDomain(d *domain.Account) *subAccountDomain.SubAccount {
	res := &subAccountDomain.SubAccount{
		AccountId:   d.AccountId,
		AccountName: d.AccountName,
		StartDate:   d.AccountStartDate,
		EndDate:     d.AccountEndDate,
	}
	return res
}

func (s *SubAccountDto) ToDomain(d *subAccountDomain.SubAccount) *domain.Account {
	return &domain.Account{
		AccountId:        d.AccountId,
		AccountName:      d.AccountName,
		AccountStartDate: d.StartDate,
		AccountEndDate:   d.EndDate,
	}
}

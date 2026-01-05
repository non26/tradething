package dto

import (
	"tradething/app/bn/future/BFF/trade/domain"
	subaccountDomain "tradething/app/bn/future/sub_account/domain"
)

type SubAccount struct {
	AccountId   string
	AccountName string
	StartDate   string
	EndDate     string
	IsExpired   bool
}

func NewSubAccountDto() *SubAccount {
	return &SubAccount{}
}

func (s *SubAccount) ToDomain(d *subaccountDomain.SubAccount) *domain.Order {
	if d == nil {
		return nil
	}
	return &domain.Order{
		AccountId: s.AccountId,
	}
}

func (s *SubAccount) FromDomain(d *domain.Order) *subaccountDomain.SubAccount {
	return &subaccountDomain.SubAccount{
		AccountId: d.AccountId,
	}
}

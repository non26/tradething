package res

import "tradething/app/bn/future/BFF/account_management/domain"

type GetRes struct {
	AccountId        string `json:"accountId"`
	AccountName      string `json:"accountName"`
	AccountStartDate string `json:"accountStartDate"`
	AccountEndDate   string `json:"accountEndDate"`
}

func (r *GetRes) FromDomain(account *domain.Account) {
	r.AccountId = account.AccountId
	r.AccountName = account.AccountName
	r.AccountStartDate = account.AccountStartDate
	r.AccountEndDate = account.AccountEndDate
}

type GetAllRes struct {
	Accounts []*GetRes `json:"accounts"`
}

func (r *GetAllRes) FromDomain(accounts []*domain.Account) {
	r.Accounts = make([]*GetRes, len(accounts))
	for i, account := range accounts {
		r.Accounts[i] = &GetRes{}
		r.Accounts[i].FromDomain(account)
	}
}

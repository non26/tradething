package res

import "tradething/app/bn/future/sub_account/domain"

type GetSubAccountRes struct {
	AccountId   string `json:"account_id"`
	AccountName string `json:"account_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsExpired   bool   `json:"is_expired"`
}

func (r *GetSubAccountRes) FromDomain(subAccount *domain.SubAccount) *GetSubAccountRes {
	if subAccount == nil {
		return nil
	}
	res := &GetSubAccountRes{
		AccountId:   subAccount.AccountId,
		AccountName: subAccount.AccountName,
		StartDate:   subAccount.StartDate,
		EndDate:     subAccount.EndDate,
		IsExpired:   subAccount.IsExpired,
	}
	return res
}

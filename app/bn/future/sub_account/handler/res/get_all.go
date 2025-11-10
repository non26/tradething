package res

import "tradething/app/bn/future/sub_account/domain"

type GetAllSubAccountRes struct {
	SubAccounts []*GetSubAccountRes `json:"sub_accounts"`
}

func (r *GetAllSubAccountRes) FromDomain(subAccounts []*domain.SubAccount) *GetAllSubAccountRes {
	for _, subAccount := range subAccounts {
		singleSubAccount := &GetSubAccountRes{}
		r.SubAccounts = append(r.SubAccounts, singleSubAccount.FromDomain(subAccount))
	}
	return r
}

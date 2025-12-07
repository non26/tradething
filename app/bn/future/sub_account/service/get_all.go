package service

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
)

func (s *subAccountService) GetAllSubAccount(ctx context.Context) ([]*domain.SubAccount, error) {
	subAccountsDB, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if subAccountsDB == nil {
		return nil, nil
	}
	subAccounts := make([]*domain.SubAccount, len(subAccountsDB))
	for i, subAccountDB := range subAccountsDB {
		subAccounts[i] = &domain.SubAccount{
			AccountId:   subAccountDB.AccountId,
			AccountName: subAccountDB.AccountName,
			StartDate:   subAccountDB.StartDate,
			EndDate:     subAccountDB.EndDate,
		}
		subAccounts[i].SetExpired()
	}
	return subAccounts, nil
}

package service

import (
	"context"
	"tradething/app/bn/future/BFF/account_management/domain"
)

func (s *accountManagementService) GetAllSubAccount(ctx context.Context) ([]*domain.Account, error) {
	subAccounts, err := s.subAccountExternalService.GetSubAccount().GetAllSubAccount(ctx)
	if err != nil {
		return nil, err
	}
	accounts := make([]*domain.Account, len(subAccounts))
	for i, subAccount := range subAccounts {
		accounts[i] = &domain.Account{
			AccountId:        subAccount.AccountId,
			AccountName:      subAccount.AccountName,
			AccountStartDate: subAccount.StartDate,
			AccountEndDate:   subAccount.EndDate,
		}
	}
	return accounts, nil
}

package service

import (
	"context"
	"tradething/app/bn/future/BFF/account_management/domain"
)

func (s *accountManagementService) GetAllSubAccount(ctx context.Context) ([]*domain.Account, error) {
	subAccounts, err := s.subAccountExternalService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return subAccounts, nil
}

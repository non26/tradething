package service

import (
	"context"
	"errors"
	"tradething/app/bn/future/BFF/account_management/domain"
)

func (s *accountManagementService) UpdateSubAccount(ctx context.Context, account *domain.Account) error {

	existingAccount, err := s.subAccountExternalService.Get(ctx, account.AccountId)
	if err != nil {
		return err
	}
	if existingAccount == nil {
		return errors.New("update account not found")
	}

	err = s.subAccountExternalService.Upsert(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

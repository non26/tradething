package service

import (
	"context"
	"tradething/app/bn/future/BFF/account_management/domain"
)

func (s *accountManagementService) InsertSubAccount(ctx context.Context, account *domain.Account) error {
	// TODO: Check if the account already exists
	err := s.subAccountExternalService.Upsert(ctx, account)
	if err != nil {
		return err
	}
	return nil
}

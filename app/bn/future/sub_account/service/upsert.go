package service

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
)

func (s *subAccountService) UpsertSubAccount(ctx context.Context, sub_account *domain.SubAccount) error {
	err := s.repository.Upsert(ctx, sub_account)
	if err != nil {
		return err
	}
	return nil
}

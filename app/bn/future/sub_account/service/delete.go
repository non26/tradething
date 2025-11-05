package service

import (
	"context"
)

func (s *subAccountService) DeleteSubAccount(ctx context.Context, accountId string) error {
	err := s.repository.Delete(ctx, accountId)
	if err != nil {
		return err
	}
	return nil
}

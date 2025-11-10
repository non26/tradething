package service

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
)

func (s *subAccountService) GetSubAccount(ctx context.Context, accountId string) (*domain.SubAccount, error) {
	subAccountDB, err := s.repository.Get(ctx, accountId)
	if err != nil {
		return nil, err
	}
	if subAccountDB == nil {
		return nil, nil
	}
	subAccount := &domain.SubAccount{
		AccountId:   subAccountDB.AccountId,
		AccountName: subAccountDB.AccountName,
		StartDate:   subAccountDB.StartDate,
		EndDate:     subAccountDB.EndDate,
	}
	return subAccount, nil
}

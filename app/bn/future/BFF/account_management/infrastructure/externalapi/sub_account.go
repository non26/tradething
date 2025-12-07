package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/account_management/domain"
	"tradething/app/bn/future/BFF/account_management/infrastructure/externalapi/dto"
	"tradething/app/bn/future/sub_account/service"
)

type subAccountExternalService struct {
	subAccountService service.ISubAccountService
}

type ISubAccountExternalService interface {
	GetAll(ctx context.Context) ([]*domain.Account, error)
	Upsert(ctx context.Context, account *domain.Account) error
	Delete(ctx context.Context, accountId string) error
	Get(ctx context.Context, accountId string) (*domain.Account, error)
}

func NewSubAccountExternalService(subAccountService service.ISubAccountService) ISubAccountExternalService {
	return &subAccountExternalService{subAccountService: subAccountService}
}

func (s *subAccountExternalService) GetAll(ctx context.Context) ([]*domain.Account, error) {
	subAccounts, err := s.subAccountService.GetAllSubAccount(ctx)
	if err != nil {
		return nil, err
	}

	if len(subAccounts) == 0 {
		return nil, nil
	}

	subAccountDto := dto.NewSubAccountDto()
	subAccountDomains := make([]*domain.Account, len(subAccounts))
	for i, subAccount := range subAccounts {
		subAccountDomains[i] = subAccountDto.ToDomain(subAccount)
	}
	return subAccountDomains, nil
}

func (s *subAccountExternalService) Upsert(ctx context.Context, account *domain.Account) error {
	subAccountDto := dto.NewSubAccountDto()
	subAccountDomain := subAccountDto.FromDomain(account)
	err := s.subAccountService.UpsertSubAccount(ctx, subAccountDomain)
	if err != nil {
		return err
	}
	return nil
}

func (s *subAccountExternalService) Delete(ctx context.Context, accountId string) error {
	err := s.subAccountService.DeleteSubAccount(ctx, accountId)
	if err != nil {
		return err
	}
	return nil
}

func (s *subAccountExternalService) Get(ctx context.Context, accountId string) (*domain.Account, error) {
	subAccount, err := s.subAccountService.GetSubAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}

	if subAccount == nil {
		return nil, nil
	}

	subAccountDto := dto.NewSubAccountDto()
	return subAccountDto.ToDomain(subAccount), nil
}

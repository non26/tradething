package externalapi

import (
	"context"
	"tradething/app/bn/future/BFF/trade/domain"
	"tradething/app/bn/future/BFF/trade/infrastructure/externalapi/sub_account/dto"
	"tradething/app/bn/future/sub_account/service"

	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type subAccountExternalService struct {
	subAccountService service.ISubAccountService
}

type ISubAccountExternalService interface {
	Get(ctx context.Context, accountId string) (*domain.Order, error)
	GetAll(ctx context.Context) ([]*domain.Order, error)
	Upsert(ctx context.Context, account *domain.Order) error
	Delete(ctx context.Context, accountId string) error
}

func NewSubAccountExternalService(subAccountService service.ISubAccountService) ISubAccountExternalService {
	return &subAccountExternalService{subAccountService: subAccountService}
}

func (s *subAccountExternalService) Get(ctx context.Context, accountId string) (*domain.Order, error) {
	subAccount, err := s.subAccountService.GetSubAccount(ctx, accountId)
	if err != nil {
		if err.Error() == appresponse.SubAccountNotRegisteredErrorCode {
			return nil, nil
		}
		return nil, err
	}
	subAccountDto := dto.NewSubAccountDto()
	res := subAccountDto.ToDomain(subAccount)
	return res, nil
}

func (s *subAccountExternalService) GetAll(ctx context.Context) ([]*domain.Order, error) {
	subAccounts, err := s.subAccountService.GetAllSubAccount(ctx)
	if err != nil {
		return nil, err
	}
	subAccountDto := dto.NewSubAccountDto()
	subAccountDomains := make([]*domain.Order, len(subAccounts))
	for i, subAccount := range subAccounts {
		subAccountDomains[i] = subAccountDto.ToDomain(subAccount)
	}
	return subAccountDomains, nil
}

func (s *subAccountExternalService) Upsert(ctx context.Context, account *domain.Order) error {
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

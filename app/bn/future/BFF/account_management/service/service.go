package service

import (
	"context"
	"tradething/app/bn/future/BFF/account_management/domain"
	"tradething/app/bn/future/BFF/account_management/infrastructure/externalapi"
)

type IAccountManagementService interface {
	GetAllSubAccount(ctx context.Context) ([]*domain.Account, error)
	InsertSubAccount(ctx context.Context, account *domain.Account) error
	UpdateSubAccount(ctx context.Context, account *domain.Account) error
	DeleteSubAccount(ctx context.Context, accountId string) error
}

type accountManagementService struct {
	subAccountExternalService externalapi.ISubAccountExternalService
}

func NewAccountManagementService(subAccountExternalService externalapi.ISubAccountExternalService) IAccountManagementService {
	return &accountManagementService{subAccountExternalService: subAccountExternalService}
}

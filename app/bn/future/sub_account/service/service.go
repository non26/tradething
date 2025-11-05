package service

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
	itfdb "tradething/app/bn/future/sub_account/infrastructure/db"
)

type ISubAccountService interface {
	GetSubAccount(ctx context.Context, accountId string) (*domain.SubAccount, error)
	GetAllSubAccount(ctx context.Context) ([]*domain.SubAccount, error)
	UpsertSubAccount(ctx context.Context, sub_account *domain.SubAccount) error
	DeleteSubAccount(ctx context.Context, accountId string) error
}

type subAccountService struct {
	repository itfdb.IBnFtRegisterAccountRepository
}

func NewSubAccountService(registerAccountRepository itfdb.IBnFtRegisterAccountRepository) ISubAccountService {
	return &subAccountService{repository: registerAccountRepository}
}

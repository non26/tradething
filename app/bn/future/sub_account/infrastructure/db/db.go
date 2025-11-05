package db

import (
	"context"
	"tradething/app/bn/future/sub_account/domain"
	models "tradething/app/bn/future/sub_account/infrastructure/db/models"
)

type IBnFtRegisterAccountRepository interface {
	Get(ctx context.Context, accountId string) (*models.BnFtRegisterAccount, error)
	GetAll(ctx context.Context) ([]*models.BnFtRegisterAccount, error)
	Upsert(ctx context.Context, sub_account *domain.SubAccount) error
	Delete(ctx context.Context, accountId string) error
}

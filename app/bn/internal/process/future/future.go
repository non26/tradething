package process

import (
	"context"
	response "tradething/app/bn/internal/handlers/future/res"
	domain "tradething/app/bn/internal/process/future/domain"
)

type IFuture interface {
	PlaceOrder(ctx context.Context, position domain.Position) error
	ClosePositionByClientIds(ctx context.Context, clientIds []string) (*response.CloseByClientIds, error)
}

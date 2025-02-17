package process

import (
	"context"
	response "tradething/app/bn/handlers/future/res"
	domain "tradething/app/bn/process/future/domain"
)

type IFuture interface {
	PlaceOrder(ctx context.Context, position domain.Position) (*response.Position, error)
	ClosePositionByClientIds(ctx context.Context, clientIds []string) (*response.CloseByClientIds, error)
	MultiplePosition(ctx context.Context, positions []domain.Position) (*response.MultiplePosition, error)
}

package service

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
)

func (s *BnFtAccumulationService) Upsert(ctx context.Context, accumulation *domain.Accumulation) error {
	return s.repository.Upsert(ctx, accumulation)
}

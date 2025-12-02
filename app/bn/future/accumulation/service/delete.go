package service

import (
	"context"
)

func (s *BnFtAccumulationService) Delete(ctx context.Context, clientId string) error {
	return s.repository.Delete(ctx, clientId)
}

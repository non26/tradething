package service

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
)

func (s *BnFtAccumulationService) Get(ctx context.Context, clientId string) (*domain.Accumulation, error) {
	accumulationDB, err := s.repository.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}
	if accumulationDB == nil {
		return nil, nil
	}
	accumulation := &domain.Accumulation{
		AccumID:      accumulationDB.AccumID,
		ClientID:     accumulationDB.ClientID,
		MaxAccum:     accumulationDB.MaxAccum,
		PresentAccum: accumulationDB.PresentAccum,
	}
	return accumulation, nil
}

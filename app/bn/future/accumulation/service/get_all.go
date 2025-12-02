package service

import (
	"context"
	"tradething/app/bn/future/accumulation/domain"
)

func (s *BnFtAccumulationService) GetAll(ctx context.Context) ([]*domain.Accumulation, error) {
	accumulationsDB, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if accumulationsDB == nil {
		return []*domain.Accumulation{}, nil
	}
	accumulations := make([]*domain.Accumulation, len(accumulationsDB))
	for i, accumulationDB := range accumulationsDB {
		accumulations[i] = &domain.Accumulation{
			AccumID:      accumulationDB.AccumID,
			ClientID:     accumulationDB.ClientID,
			MaxAccum:     accumulationDB.MaxAccum,
			PresentAccum: accumulationDB.PresentAccum,
		}
	}
	return accumulations, nil
}

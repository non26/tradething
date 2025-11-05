package service

import (
	"context"
	"tradething/app/bn/future/position_history/domain"
)

func (s *service) GetPositionHistory(ctx context.Context, clientId string) (*domain.BnFtHistory, error) {
	historyDB, err := s.repository.Get(ctx, clientId)
	if err != nil {
		return nil, err
	}

	history := &domain.BnFtHistory{
		ClientId:     historyDB.ClientId,
		Symbol:       historyDB.Symbol,
		PositionSide: historyDB.PositionSide,
		CreatedAt:    historyDB.CreatedAt,
	}

	return history, nil
}

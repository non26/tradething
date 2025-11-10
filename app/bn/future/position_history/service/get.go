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

	if historyDB == nil {
		return nil, nil
	}

	history := &domain.BnFtHistory{
		ClientId:     historyDB.ClientId,
		Symbol:       historyDB.Symbol,
		PositionSide: historyDB.PositionSide,
		AccountId:    historyDB.AccountId,
	}

	return history, nil
}

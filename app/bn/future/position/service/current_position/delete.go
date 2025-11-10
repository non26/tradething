package service

import "context"

func (s *currentPositionService) Delete(ctx context.Context, symbol string, accountId string, positionSide string) error {
	err := s.repository.Delete(ctx, symbol, accountId, positionSide)
	if err != nil {
		return err
	}
	return nil
}

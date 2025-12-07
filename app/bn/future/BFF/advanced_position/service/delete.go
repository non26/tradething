package service

import "context"

func (s *advancedPositionService) Delete(ctx context.Context, clientId string) error {
	err := s.advancedService.Delete(ctx, clientId)
	if err != nil {
		return err
	}
	return nil
}

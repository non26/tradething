package service

import "context"

func (s *accountManagementService) DeleteSubAccount(ctx context.Context, accountId string) error {
	err := s.subAccountExternalService.Delete(ctx, accountId)
	if err != nil {
		return err
	}
	return nil
}

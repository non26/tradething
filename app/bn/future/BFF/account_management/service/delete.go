package service

import "context"

func (s *accountManagementService) DeleteSubAccount(ctx context.Context, accountId string) error {
	err := s.subAccountExternalService.GetSubAccount().DeleteSubAccount(ctx, accountId)
	if err != nil {
		return err
	}
	return nil
}

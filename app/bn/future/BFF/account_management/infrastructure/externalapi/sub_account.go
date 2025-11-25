package externalapi

import "tradething/app/bn/future/sub_account/service"

type subAccountExternalService struct {
	subAccountService service.ISubAccountService
}

type ISubAccountExternalService interface {
	GetSubAccount() service.ISubAccountService
}

func NewSubAccountExternalService(subAccountService service.ISubAccountService) ISubAccountExternalService {
	return &subAccountExternalService{subAccountService: subAccountService}
}

func (s *subAccountExternalService) GetSubAccount() service.ISubAccountService {
	return s.subAccountService
}

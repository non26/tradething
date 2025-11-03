package service

type IService interface{}

type service struct {
}

func NewService() IService {
	return &service{}
}

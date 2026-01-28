package req

import "tradething/app/bn/future/BFF/trade/domain"

type CloseOrderByIdReq struct {
	ClientId string `json:"client_id"`
}

func (r *CloseOrderByIdReq) ToDomain() *domain.Order {
	return &domain.Order{
		ClientId: r.ClientId,
	}
}

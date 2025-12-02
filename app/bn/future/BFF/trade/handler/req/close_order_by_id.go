package req

import "tradething/app/bn/future/BFF/trade/domain"

type CloseOrderByIdReq struct {
	ID string `json:"id"`
}

func (r *CloseOrderByIdReq) ToDomain() *domain.Order {
	return &domain.Order{
		ClientId: r.ID,
	}
}

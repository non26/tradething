package req

import "tradething/app/bn/future/BFF/advanced_position/domain"

type GetReq struct {
	ClientId string `json:"client_id" binding:"required"`
}

func (r *GetReq) ToDomain() *domain.AdvancedPosition {
	return &domain.AdvancedPosition{
		ClientId: r.ClientId,
	}
}

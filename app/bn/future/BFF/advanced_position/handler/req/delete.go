package req

import "tradething/app/bn/future/BFF/advanced_position/domain"

type DeleteReq struct {
	ClientId string `json:"client_id" binding:"required"`
}

func (r *DeleteReq) ToDomain() *domain.AdvancedPosition {
	return &domain.AdvancedPosition{
		ClientId: r.ClientId,
	}
}

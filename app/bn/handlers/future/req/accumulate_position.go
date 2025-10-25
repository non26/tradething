package req

import "tradething/app/bn/process/future/domain"

type AccumulatePositionReq struct {
	ParentClientId        string `json:"parent_client_id"`
	MaxAccumulatePosition string `json:"max_accumulate_position"`
	AccumulateAmountB     string `json:"accumulate_amount_b"`
}

func (a *AccumulatePositionReq) ToDomain() *domain.Position {
	position := domain.NewPosition()
	position.SetClientId(a.ParentClientId)
	position.SetMaxAccumulatePosition(a.MaxAccumulatePosition)
	position.SetEntryQuantity(a.AccumulateAmountB)
	return position
}

package req

import (
	"time"
	"tradething/app/bn/future/sub_account/domain"
)

type UpsertReq struct {
	AccountId   string `json:"accountId" binding:"required"`
	AccountName string `json:"accountName" binding:"required"`
	StartDate   string `json:"startDate" binding:"required"`
	EndDate     string `json:"endDate" binding:"required"`
}

func (r *UpsertReq) ToDomain() *domain.SubAccount {
	return &domain.SubAccount{
		AccountId:   r.AccountId,
		AccountName: r.AccountName,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
	}
}

func (r *UpsertReq) Validate() error {
	_, err := time.Parse(time.RFC3339, r.StartDate)
	if err != nil {
		return err
	}
	_, err = time.Parse(time.RFC3339, r.EndDate)
	if err != nil {
		return err
	}
	return nil
}

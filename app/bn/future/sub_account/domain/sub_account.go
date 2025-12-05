package domain

import "time"

type SubAccount struct {
	AccountId   string
	AccountName string
	StartDate   string
	EndDate     string
	IsExpired   bool
}

func (s *SubAccount) IsAccountExpired() bool {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return s.EndDate < time.Now().In(loc).Format("2006-01-02")
}

func (s *SubAccount) SetExpired() {
	s.IsExpired = s.IsAccountExpired()
}

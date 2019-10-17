package models

import "time"

type Rebate struct {
	Id      int
	Title   string
	TallyMoney   string
	Rebate  string
	StartDate time.Time
	EndDate time.Time
	Created time.Time
	Updated time.Time
}

func (m *Rebate) TableName() string {
	return TableName("rebate")
}
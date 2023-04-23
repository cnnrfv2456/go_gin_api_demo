package models

import "time"

type Member struct {
	Id        int64
	Name      string
	Account   string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (t *Member) TableName() string {
	return "member"
}

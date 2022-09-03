package domain

import "time"

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	User       User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

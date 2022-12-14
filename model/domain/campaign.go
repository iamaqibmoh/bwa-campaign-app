package domain

import "time"

type Campaign struct {
	Id             int
	UserId         int
	Name           string
	Summary        string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmount     int
	CurrentAmount  int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
	User           User
}

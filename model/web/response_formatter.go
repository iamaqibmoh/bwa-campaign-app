package web

import "time"

type UserResponseFormatter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type CampaignResponseFormatterBasic struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	ImageUrl      string `json:"image_url"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
	Slug          string `json:"slug"`
}

type CampaignResponseFormatter struct {
	Id            int                      `json:"id"`
	Name          string                   `json:"name"`
	Summary       string                   `json:"summary"`
	Description   string                   `json:"description"`
	ImageUrl      string                   `json:"image_url"`
	GoalAmount    int                      `json:"goal_amount"`
	CurrentAmount int                      `json:"current_amount"`
	UserId        int                      `json:"user_id"`
	Slug          string                   `json:"slug"`
	Perks         []string                 `json:"perks"`
	User          CampaignUserFormatter    `json:"user"`
	Images        []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

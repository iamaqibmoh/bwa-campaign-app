package web

type UserResponseFormatter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type CampaignResponseFormatter struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	ImageUrl      string `json:"image_url"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
	Slug          string `json:"slug"`
}

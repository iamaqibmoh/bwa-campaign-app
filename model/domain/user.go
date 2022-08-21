package domain

import "time"

type User struct {
	Id           int
	Name         string
	Occupation   string
	Email        string
	PasswordHash string
	Avatar       string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

package repository

import "BWA-CAMPAIGN-APP/model/domain"

type UserRepository interface {
	Save(user domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id int) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}

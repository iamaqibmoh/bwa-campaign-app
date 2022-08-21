package repository

import "BWA-CAMPAIGN-APP/model/domain"

type Repository interface {
	Save(user domain.User) (domain.User, error)
}

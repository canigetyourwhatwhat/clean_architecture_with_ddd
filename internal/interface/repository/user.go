package repository

import "clean_architecture_with_ddd/internal/entity"

type UserRepository interface {
	GetUserByUsername(username string) (*entity.User, error)
	GetTaxFromUserByTaxId(userId int) (float32, error)

	CreateUser(user *entity.User) error
}

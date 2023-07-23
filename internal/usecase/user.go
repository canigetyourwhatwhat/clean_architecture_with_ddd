package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"database/sql"
	"errors"
)

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{
		repo: repo,
	}
}

type UserService interface {
	CreateUser(lastName string, firstName string, password string, username string) error
}

func (u *userService) CreateUser(lastName string, firstName string, password string, username string) error {

	// validate if the user is already created
	if user, err := u.repo.GetUserByUsername(username); err != nil && err != sql.ErrNoRows {
		return err
	} else if user != nil {
		return errors.New("user is already created with this username")
	}

	user := &entity.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}
	if err := user.SetHashedPassword(password); err != nil {
		return err
	}

	if err := u.repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

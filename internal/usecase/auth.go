package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"time"
)

type authService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) AuthUsecase {
	return &authService{
		repo: repo,
	}
}

type AuthUsecase interface {
	Login(username string, password string) (sessionID string, err error)
}

func (a *authService) Login(username string, password string) (string, error) {

	// validate if the username exists
	user, err := a.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// validate password
	if err = user.ComparePassword(password); err != nil {
		return "", err
	}

	// store the session information
	session := &entity.Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 12),
	}
	if err = session.SetSessionID(); err != nil {
		return "", err
	}
	if err = a.repo.CreateSession(session); err != nil {
		return "", err
	}

	return session.ID, nil
}

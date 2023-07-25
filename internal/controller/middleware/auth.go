package middleware

import (
	"clean_architecture_with_ddd/internal/interface/repository"
	"errors"
	"github.com/labstack/echo/v4"
	"time"
)

type auth struct {
	repo repository.Repository
}

func NewAuth(repo repository.Repository) Auth {
	return &auth{
		repo: repo,
	}
}

type Auth interface {
	GetSession(c echo.Context) (int, error)
}

func (a *auth) GetSession(c echo.Context) (int, error) {
	sessionKey := c.Request().Header.Get("session")
	if sessionKey == "" {
		return -1, errors.New("session is missing")
	}
	session, err := a.repo.GetSessionById(sessionKey)
	if err != nil {
		return -1, errors.New("session is not valid")
	}

	if session.ExpiresAt.Before(time.Now()) {
		return -1, errors.New("session is already expired")
	}
	return session.UserID, nil
}

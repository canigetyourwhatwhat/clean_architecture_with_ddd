package repository

import "clean_architecture_with_ddd/internal/entity"

type SessionRepository interface {
	CreateOrUpdateSession(session *entity.Session) error
	GetSessionById(id string) (*entity.Session, error)
}

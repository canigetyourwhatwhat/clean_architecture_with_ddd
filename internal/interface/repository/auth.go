package repository

import "clean_architecture_with_ddd/internal/entity"

type SessionRepository interface {
	CreateSession(session *entity.Session) error
}

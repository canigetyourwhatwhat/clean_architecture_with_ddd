package repository

import "clean_architecture_with_ddd/internal/entity"

func (r Repo) CreateOrUpdateSession(session *entity.Session) error {
	query := `
	INSERT INTO sessions
	(
	 	id,
		userId,
		expiresAt
	)
	VALUES
	(
	 	:id,
		:userId,
		:expiresAt
	)
	ON DUPLICATE KEY UPDATE id = :id, userId = :userId, expiresAt = :expiresAt;
	`
	_, err := r.DB.NamedExec(query, session)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) GetSessionById(id string) (*entity.Session, error) {
	var s entity.Session
	err := r.DB.Get(s, "select * from sessions where id = ?", id)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

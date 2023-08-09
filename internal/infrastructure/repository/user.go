package repository

import (
	"clean_architecture_with_ddd/internal/entity"
)

func (r Repo) CreateUser(user *entity.User) error {
	query := `
	INSERT INTO users
	(firstName, lastName, username, password)
	VALUES
	(
		:firstName,
	 	:lastName,
	 	:username,
	 	:password
	)
	`
	_, err := r.DB.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Get(&user, "select * from users where username = ?", username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r Repo) GetTaxFromUserByTaxId(userId int) (tax int, err error) {
	err = r.DB.Get(&tax, "select taxId from users where id = ?", userId)
	if err != nil {
		return 0, err
	}
	return tax, nil
}

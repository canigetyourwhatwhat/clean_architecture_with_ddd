package inputPort

import "errors"

func CreateUser(lastName string, firstName string, password string, username string) error {
	if lastName == "" || firstName == "" || password == "" || username == "" {
		return errors.New("input is empty")
	}
	if len(password) <= 8 {
		return errors.New("password is too weak")
	}
	return nil
}

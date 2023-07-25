package inputPort

import "errors"

func Login(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("input is empty")
	}
	return nil
}

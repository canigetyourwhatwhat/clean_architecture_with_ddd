package inputPort

import (
	"errors"
	"strconv"
)

func ID(id string) error {
	if id == "" {
		return errors.New("product ID is empty")
	}
	return nil
}

func IntID(id string) (int, error) {
	if id == "" {
		return -1, errors.New("product ID is empty")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return -1, errors.New("ID is not integer")
	}
	return intId, nil
}

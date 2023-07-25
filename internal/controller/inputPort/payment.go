package inputPort

import (
	"clean_architecture_with_ddd/internal/entity"
	"errors"
)

func CreatePayment(method entity.PaymentMethod) error {
	if method != entity.Cash || method != entity.Card {
		return errors.New("input is invalid")
	}
	return nil
}

package inputPort

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"errors"
)

func CartItem(productCode string, quantity int) error {
	if productCode == "" || quantity < 1 {
		return errors.New("input is invalid")
	}
	return nil
}

func ListCartItem(productInfo request.ListCartItem) error {
	if len(productInfo.CartItems) == 0 {
		return errors.New("input is empty")
	}
	return nil
}

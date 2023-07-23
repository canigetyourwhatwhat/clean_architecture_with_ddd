package repository

import "clean_architecture_with_ddd/internal/entity"

type CartRepository interface {
	GetCartByStatusAndUserId(status entity.CartStatus, userID int) (*entity.Cart, error)
	GetCartById(id int) (*entity.Cart, error)

	UpdateCart(cart *entity.Cart) error
	CreateCart(cart *entity.Cart) error
}

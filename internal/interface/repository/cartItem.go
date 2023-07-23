package repository

import "clean_architecture_with_ddd/internal/entity"

type CartItemRepository interface {
	GetCartItemByCodeAndCartId(code string, cartID int) (*entity.CartItem, error)
	ListCartItemByCartId(id int) ([]entity.CartItem, error)

	DeleteCartItemById(id int) error
	DeleteCartItemByCartId(id int) error
	CreateCartItem(ci *entity.CartItem) error
}

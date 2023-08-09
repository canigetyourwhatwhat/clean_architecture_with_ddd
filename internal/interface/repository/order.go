package repository

import "clean_architecture_with_ddd/internal/entity"

type OrderRepository interface {
	GetOrderById(id int) (*entity.Order, error)
	CreateOrder(order entity.Order) error
}

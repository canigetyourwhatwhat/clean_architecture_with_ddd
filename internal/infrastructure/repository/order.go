package repository

import "clean_architecture_with_ddd/internal/entity"

type OrderRepository interface {
	GetOrderById(id int) (*entity.Order, error)
	CreateOrder(order entity.Order) error
}

func (r Repo) GetOrderById(id int) (*entity.Order, error) {
	var o entity.Order
	if err := r.DB.Get(&o, "select * from orders where id = ?", id); err != nil {
		return nil, err
	}
	return &o, nil
}

func (r Repo) CreateOrder(order entity.Order) error {
	query := `INSERT INTO orders (userId, cartId, paymentId) VALUES (:userId, :cartId, :paymentId)`
	_, err := r.DB.NamedExec(query, order)
	if err != nil {
		return err
	}
	return nil
}

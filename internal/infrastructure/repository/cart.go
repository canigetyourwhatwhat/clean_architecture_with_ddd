package repository

import "clean_architecture_with_ddd/internal/entity"

func (r Repo) ListCartsByStatusAndUserId(status entity.CartStatus, userId int) ([]*entity.Cart, error) {
	var c []*entity.Cart
	if err := r.DB.Select(&c, "select * from carts where userId = ? and status = ?", userId, status); err != nil {
		return nil, err
	}
	return c, nil
}

func (r Repo) GetCartById(id int) (*entity.Cart, error) {
	var c entity.Cart
	if err := r.DB.Get(&c, "select * from carts where id = ?", id); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r Repo) UpdateCart(cart *entity.Cart) error {
	query := `UPDATE carts set netPrice = :netPrice, taxPrice = :taxPrice, totalPrice = :totalPrice, status = :status where id = :id`
	_, err := r.DB.NamedExec(query, cart)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) CreateCart(cart *entity.Cart) error {
	query := `INSERT INTO carts (userId) VALUES (:userId)`
	_, err := r.DB.NamedExec(query, cart)
	if err != nil {
		return err
	}
	return nil
}

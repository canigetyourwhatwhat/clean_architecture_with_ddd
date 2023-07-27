package repository

import "clean_architecture_with_ddd/internal/entity"

type CartItemRepository interface {
	GetCartItemByCodeAndCartId(code string, cartID int) (*entity.CartItem, error)
	ListCartItemByCartId(id int) ([]entity.CartItem, error)

	DeleteCartItemById(id int) error
	DeleteCartItemByCartId(id int) error
	CreateCartItem(ci *entity.CartItem) error
}

func (r Repo) GetCartItemByCodeAndCartId(code string, cartID int) (ci *entity.CartItem, err error) {
	query := `select * from cartItems where cartId = ? and productCode = ?`
	err = r.DB.Get(&ci, query, cartID, code)
	if err != nil {
		return nil, err
	}
	return ci, nil
}

func (r Repo) ListCartItemByCartId(id int) (ci []entity.CartItem, err error) {
	query := `select * from cartItems where cartId = ?`
	err = r.DB.Select(&ci, query, id)
	if err != nil {
		return nil, err
	}
	return ci, nil
}

func (r Repo) DeleteCartItemById(id int) error {
	query := `delete from cartItems where id = ?`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) DeleteCartItemByCartId(id int) error {
	query := `delete from cartItems where cartId = ?`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) CreateCartItem(ci *entity.CartItem) error {
	query := `INSERT INTO cartItems (productCode, cartId, quantity, totalPrice, taxPrice, netPrice) VALUES (:productCode, :cartId, :quantity, :totalPrice, :taxPrice, :netPrice)`
	_, err := r.DB.NamedExec(query, *ci)
	if err != nil {
		return err
	}
	return nil
}

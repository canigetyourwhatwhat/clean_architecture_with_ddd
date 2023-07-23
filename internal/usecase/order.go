package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"database/sql"
	"errors"
)

type orderService struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) OrderService {
	return &orderService{
		repo: repo,
	}
}

type OrderService interface {
	GetOrder(userID int, orderId int) (*entity.Order, error)
}

func (o *orderService) GetOrder(userID int, orderId int) (*entity.Order, error) {

	order, err := o.repo.GetOrderById(orderId)
	if err == sql.ErrNoRows {
		return nil, errors.New("order doesn't exist")
	}
	if err != nil {
		return nil, err
	}

	payment, err := o.repo.GetPaymentById(order.PaymentId)
	if err != nil {
		return nil, err
	}

	cart, err := o.repo.GetCartById(order.CartId)
	if err != nil {
		return nil, err
	}

	cart.CartItems, err = o.repo.ListCartItemByCartId(cart.ID)
	if err != nil {
		return nil, err
	}

	response := &entity.Order{
		ID:      orderId,
		UserId:  userID,
		Payment: *payment,
		Cart:    *cart,
	}

	return response, nil
}

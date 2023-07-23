package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"database/sql"
	"errors"
)

type paymentService struct {
	repo repository.Repository
}

func NewPaymentService(repo repository.Repository) PaymentService {
	return &paymentService{
		repo: repo,
	}
}

type PaymentService interface {
	CompleteShopping(userID int, method int) error
}

func (p *paymentService) CompleteShopping(userID int, method int) error {
	// get the current shopping cart
	cart, err := p.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
	if err == sql.ErrNoRows {
		return errors.New("there is no shopping cart")
	}
	if err != nil {
		return err
	}

	// create payment and store it
	payment := entity.Payment{
		UserId: userID,
		Amount: cart.TotalPrice,
		Method: entity.PaymentMethod(method),
	}
	paymentId, err := p.repo.CreatePayment(payment)
	if err != nil {
		return err
	}

	// create order and store it
	order := entity.Order{
		UserId:    cart.UserId,
		CartId:    cart.ID,
		PaymentId: paymentId,
	}
	err = p.repo.CreateOrder(order)
	if err != nil {
		return err
	}

	// update the status of the cart to "completed"
	cart.Status = entity.Completed
	err = p.repo.UpdateCart(cart)
	if err != nil {
		return err
	}

	return nil
}

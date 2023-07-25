package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
)

type cartUsecase struct {
	repo repository.Repository
}

func NewCartService(repo repository.Repository) CartUsecase {
	return &cartUsecase{
		repo: repo,
	}
}

type CartUsecase interface {
	GetInProgressCart(userID int) (*entity.Cart, error)
}

func (c cartUsecase) GetInProgressCart(userID int) (*entity.Cart, error) {
	cart, err := c.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
	if err != nil {
		return nil, err
	}

	// put all the cartItem in this object
	cartItems, err := c.repo.ListCartItemByCartId(cart[0].ID)
	if err != nil {
		return nil, err
	}
	cart[0].CartItems = cartItems

	return cart[0], nil
}

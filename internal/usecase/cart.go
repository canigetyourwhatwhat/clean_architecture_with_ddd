package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
)

type cartService struct {
	repo repository.Repository
}

func NewCartService(repo repository.Repository) CartService {
	return &cartService{
		repo: repo,
	}
}

type CartService interface {
	GetInProgressCart(userID int) (*entity.Cart, error)
}

func (c cartService) GetInProgressCart(userID int) (*entity.Cart, error) {
	cart, err := c.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
	if err != nil {
		return nil, err
	}

	// put all the cartItem in this object
	cartItems, err := c.repo.ListCartItemByCartId(cart.ID)
	if err != nil {
		return nil, err
	}
	cart.CartItems = cartItems

	return cart, nil
}

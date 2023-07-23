package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"database/sql"
	"errors"
)

type cartItemService struct {
	repo repository.Repository
}

func NewCartItemService(repo repository.Repository) CartItemService {
	return &cartItemService{
		repo: repo,
	}
}

type CartItemService interface {
	AddItemInCart(userID int, productInfo CartItemRequest) error
	DeleteItemFromCart(userID int, productCode string) error
	UpdateItemsInCart(userID int, productInfo []CartItemRequest) error
}

func (ci cartItemService) AddItemInCart(userID int, productInfo CartItemRequest) error {
	// get the current shopping cart
	cart, err := ci.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Create cart if it doesn't exist
	if cart == nil {
		err = ci.repo.CreateCart(&entity.Cart{UserId: userID})
		if err != nil {
			return err
		}

		cart, err = ci.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
		if err != nil {
			return err
		}
	}

	// check if the product is already added
	cartItem, err := ci.repo.GetCartItemByCodeAndCartId(productInfo.ProductCode, cart.ID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// get product to calculate price
	product, err := ci.repo.GetProductByCode(productInfo.ProductCode)
	if err != nil {
		return err
	}

	// get tax rate
	taxRate, err := ci.repo.GetTaxFromUserByTaxId(userID)
	if err != nil {
		return err
	}

	// if already this product is added
	if cartItem != nil {
		cart.NetPrice = cart.NetPrice - cartItem.NetPrice + product.Price*float32(productInfo.Quantity)
		cart.TaxPrice = cart.TaxPrice - cartItem.TaxPrice + product.Price*float32(productInfo.Quantity)*taxRate
		cart.TotalPrice = cart.TotalPrice - cartItem.TotalPrice + product.Price*float32(productInfo.Quantity)*(1+taxRate)

		err = ci.repo.DeleteCartItemById(cartItem.ID)
		if err != nil {
			return err
		}
	} else {
		cart.NetPrice = cart.NetPrice + product.Price*float32(productInfo.Quantity)
		cart.TaxPrice = cart.TaxPrice + product.Price*float32(productInfo.Quantity)*taxRate
		cart.TotalPrice = cart.TotalPrice + product.Price*float32(productInfo.Quantity)*(1+taxRate)
	}

	err = ci.repo.UpdateCart(cart)
	if err != nil {
		return err
	}

	cartItem = &entity.CartItem{
		ProductCode: product.Code,
		CartId:      cart.ID,
		Quantity:    productInfo.Quantity,
		NetPrice:    cart.NetPrice,
		TaxPrice:    cart.TaxPrice,
		TotalPrice:  cart.TotalPrice,
	}

	err = ci.repo.CreateCartItem(cartItem)
	if err != nil {
		return err
	}

	return nil
}

func (ci cartItemService) DeleteItemFromCart(userID int, code string) error {

	// get the current shopping cart
	cart, err := ci.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
	if err == sql.ErrNoRows {
		return errors.New("cart doesn't exist")
	}
	if err != nil {
		return err
	}

	// if this product is already added, delete it
	cartItem, err := ci.repo.GetCartItemByCodeAndCartId(code, cart.ID)
	if err == sql.ErrNoRows {
		return errors.New("this product is not in the cart")
	}
	if err != nil {
		return err
	}
	err = ci.repo.DeleteCartItemById(cartItem.ID)
	if err != nil {
		return err
	}

	// calculate price and update cart
	cart.NetPrice = cart.NetPrice - cartItem.NetPrice
	cart.TaxPrice = cart.TaxPrice - cartItem.TaxPrice
	cart.TotalPrice = cart.TotalPrice - cartItem.TotalPrice
	err = ci.repo.UpdateCart(cart)
	if err != nil {
		return err
	}

	return nil
}

func (ci cartItemService) UpdateItemsInCart(userID int, productInfo []CartItemRequest) error {

	// get the current shopping cart
	cart, err := ci.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// create cart if it doesn't exist
	if cart == nil {
		err = ci.repo.CreateCart(&entity.Cart{UserId: userID})
		if err != nil {
			return err
		}

		cart, err = ci.repo.GetCartByStatusAndUserId(entity.InProgress, userID)
		if err != nil {
			return err
		}
	}

	// Delete all the items in the cart
	err = ci.repo.DeleteCartItemByCartId(cart.ID)

	// reset the cart info
	cart.NetPrice = 0
	cart.TaxPrice = 0
	cart.TotalPrice = 0

	// get tax rate
	taxRate, err := ci.repo.GetTaxFromUserByTaxId(userID)
	if err != nil {
		return err
	}

	// Store list of product with quantity and calculate costs for cart
	var product *entity.Product
	for i := range productInfo {
		product, err = ci.repo.GetProductByCode(productInfo[i].ProductCode)
		if err != nil {
			return err
		}

		cartItem := &entity.CartItem{
			CartId:      cart.ID,
			Quantity:    productInfo[i].Quantity,
			ProductCode: productInfo[i].ProductCode,
			NetPrice:    product.Price * float32(productInfo[i].Quantity),
			TaxPrice:    product.Price * float32(productInfo[i].Quantity) * taxRate,
			TotalPrice:  product.Price * float32(productInfo[i].Quantity) * (1 + taxRate),
		}

		err = ci.repo.CreateCartItem(cartItem)
		if err != nil {
			return err
		}

		cart.NetPrice += cartItem.NetPrice
		cart.TaxPrice += cartItem.TaxPrice
		cart.TotalPrice += cartItem.TotalPrice
	}

	err = ci.repo.UpdateCart(cart)
	if err != nil {
		return err
	}

	return nil
}

//  ----- request body ------

type CartItemRequest struct {
	ProductCode string
	Quantity    int
}

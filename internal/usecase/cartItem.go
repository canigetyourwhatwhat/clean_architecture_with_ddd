package usecase

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
	"database/sql"
	"errors"
)

type cartItemService struct {
	repo repository.Repository
}

func NewCartItemService(repo repository.Repository) CartItemUsecase {
	return &cartItemService{
		repo: repo,
	}
}

type CartItemUsecase interface {
	AddItemInCart(userID int, productInfo request.CartItem) error
	DeleteItemFromCart(userID int, productCode string) error
	UpdateItemsInCart(userID int, productInfo request.ListCartItem) error

	GetPurchasedProducts(userID int) ([]entity.CartItem, error)
}

func (ci *cartItemService) AddItemInCart(userID int, productInfo request.CartItem) error {
	// get the current shopping cart
	carts, err := ci.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Create cart if it doesn't exist
	if len(carts) == 0 {
		err = ci.repo.CreateCart(&entity.Cart{UserId: userID})
		if err != nil {
			return err
		}

		carts, err = ci.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
		if err != nil {
			return err
		}
	}

	// get current shopping cart
	cart := carts[0]

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
	taxId, err := ci.repo.GetTaxFromUserByTaxId(userID)
	if err != nil {
		return err
	}
	tax, err := ci.repo.GetTaxRateById(taxId)
	if err != nil {
		return err
	}

	// if already this product is added
	if cartItem != nil {
		cart.NetPrice = cart.NetPrice - cartItem.NetPrice + product.Price*float32(productInfo.Quantity)
		cart.TaxPrice = cart.TaxPrice - cartItem.TaxPrice + product.Price*float32(productInfo.Quantity)*tax
		cart.TotalPrice = cart.TotalPrice - cartItem.TotalPrice + product.Price*float32(productInfo.Quantity)*(1+tax)

		err = ci.repo.DeleteCartItemById(cartItem.ID)
		if err != nil {
			return err
		}
	} else {
		cart.NetPrice = cart.NetPrice + product.Price*float32(productInfo.Quantity)
		cart.TaxPrice = cart.TaxPrice + product.Price*float32(productInfo.Quantity)*tax
		cart.TotalPrice = cart.TotalPrice + product.Price*float32(productInfo.Quantity)*(1+tax)
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

func (ci *cartItemService) DeleteItemFromCart(userID int, code string) error {

	// get the current shopping cart
	carts, err := ci.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
	if len(carts) == 0 {
		return errors.New("cart doesn't exist")
	}
	if err != nil {
		return err
	}

	// if this product is already added, delete it
	cartItem, err := ci.repo.GetCartItemByCodeAndCartId(code, carts[0].ID)
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
	carts[0].NetPrice = carts[0].NetPrice - cartItem.NetPrice
	carts[0].TaxPrice = carts[0].TaxPrice - cartItem.TaxPrice
	carts[0].TotalPrice = carts[0].TotalPrice - cartItem.TotalPrice
	err = ci.repo.UpdateCart(carts[0])
	if err != nil {
		return err
	}

	return nil
}

func (ci *cartItemService) UpdateItemsInCart(userID int, productInfo request.ListCartItem) error {

	// get the current shopping cart
	carts, err := ci.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Create cart if it doesn't exist
	if len(carts) == 0 {
		err = ci.repo.CreateCart(&entity.Cart{UserId: userID})
		if err != nil {
			return err
		}

		carts, err = ci.repo.ListCartsByStatusAndUserId(entity.InProgress, userID)
		if err != nil {
			return err
		}
	}

	// get current shopping cart
	cart := carts[0]

	// Delete all the items in the cart
	err = ci.repo.DeleteCartItemByCartId(cart.ID)

	// reset the cart info
	cart.NetPrice = 0
	cart.TaxPrice = 0
	cart.TotalPrice = 0

	// get tax rate
	taxId, err := ci.repo.GetTaxFromUserByTaxId(userID)
	if err != nil {
		return err
	}
	tax, err := ci.repo.GetTaxRateById(taxId)
	if err != nil {
		return err
	}

	// Store list of product with quantity and calculate costs for cart
	var product *entity.Product
	for i := range productInfo.CartItems {
		product, err = ci.repo.GetProductByCode(productInfo.CartItems[i].ProductCode)
		if err != nil {
			return err
		}

		cartItem := &entity.CartItem{
			CartId:      cart.ID,
			Quantity:    productInfo.CartItems[i].Quantity,
			ProductCode: productInfo.CartItems[i].ProductCode,
			NetPrice:    product.Price * float32(productInfo.CartItems[i].Quantity),
			TaxPrice:    product.Price * float32(productInfo.CartItems[i].Quantity) * tax,
			TotalPrice:  product.Price * float32(productInfo.CartItems[i].Quantity) * (1 + tax),
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

func (ci *cartItemService) GetPurchasedProducts(userID int) ([]entity.CartItem, error) {

	// get the current shopping cart
	carts, err := ci.repo.ListCartsByStatusAndUserId(entity.Completed, userID)
	if len(carts) == 0 {
		return nil, errors.New("no purchased history")
	}
	if err != nil {
		return nil, err
	}

	var cartItems []entity.CartItem
	for i := range carts {
		newCartItems, err := ci.repo.ListCartItemByCartId(carts[i].ID)
		if err != nil {
			return nil, err
		}
		for i := range newCartItems {
			product, err := ci.repo.GetProductByCode(newCartItems[0].ProductCode)
			if err != nil {
				return nil, err
			}
			newCartItems[i].Product = product
		}
		cartItems = append(cartItems, newCartItems...)
	}

	return cartItems, nil
}

//  ----- request body ------

//type CartItemRequest struct {
//	ProductCode string
//	Quantity    int
//}

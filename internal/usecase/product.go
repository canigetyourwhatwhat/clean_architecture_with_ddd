package usecase

import "clean_architecture_with_ddd/internal/entity"

type ProductService struct {
}

func NewProductService() ProductService {
	return ProductService{}
}

type ProductServiceInterface interface {
	GetProduct(productID int, userID int) *entity.Product
	ListProductsByPage(page int, userID int) []*entity.Product
}

func (s *ProductService) GetProduct(productID int, userID int) *entity.Product {

	return nil
}

func (s *ProductService) ListProductsByPage(page int, userID int) []*entity.Product {
	return nil
}

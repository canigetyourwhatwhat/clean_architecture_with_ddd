package usecase

import (
	"clean_architecture_with_ddd/internal/entity"
	"clean_architecture_with_ddd/internal/interface/repository"
)

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) ProductService {
	return &productService{
		repo: repo,
	}
}

type ProductService interface {
	GetProduct(productID string) (*entity.Product, error)
	ListProductsByPage(page int) ([]*entity.Product, error)
}

func (s *productService) GetProduct(productID string) (*entity.Product, error) {
	product, err := s.repo.GetProductByCode(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) ListProductsByPage(page int) ([]*entity.Product, error) {
	products, err := s.repo.ListProductsByPageNum(page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

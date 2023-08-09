package repository

import "clean_architecture_with_ddd/internal/entity"

type ProductRepository interface {
	GetProductByCode(code string) (*entity.Product, error)
	GetProductCount() (int, error)
	ListProductsByPageNum(pageNum int, perPage int) ([]*entity.Product, error)
}

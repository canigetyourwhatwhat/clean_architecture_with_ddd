package repository

import "clean_architecture_with_ddd/internal/entity"

type ProductRepository interface {
	GetProductByCode(code string) (*entity.Product, error)
	ListProductsByPageNum(pageNum int) ([]*entity.Product, error)
}

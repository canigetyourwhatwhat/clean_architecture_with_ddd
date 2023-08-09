package presenter

import (
	"clean_architecture_with_ddd/internal/controller/entity/response"
	"clean_architecture_with_ddd/internal/entity"
)

func AllProduct(count int, products []*entity.Product) response.ListProducts {
	return response.ListProducts{
		Count:    count,
		Products: products,
	}
}

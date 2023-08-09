package response

import "clean_architecture_with_ddd/internal/entity"

type ListProducts struct {
	Count    int               `json:"count"`
	Products []*entity.Product `json:"products"`
}

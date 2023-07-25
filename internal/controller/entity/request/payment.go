package request

import "clean_architecture_with_ddd/internal/entity"

type CreatePayment struct {
	Method entity.PaymentMethod `json:"method"`
}

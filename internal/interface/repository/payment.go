package repository

import "clean_architecture_with_ddd/internal/entity"

type PaymentRepository interface {
	GetPaymentById(id int) (*entity.Payment, error)

	CreatePayment(payment entity.Payment) (int, error)
}

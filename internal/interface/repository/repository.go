package repository

type Repository interface {
	ProductRepository
	UserRepository
	SessionRepository
	CartRepository
	CartItemRepository
	PaymentRepository
	OrderRepository
}

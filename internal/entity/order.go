package entity

import "time"

type Order struct {
	ID        int `db:"id"`
	UserId    int `db:"userId"`
	CartId    int `db:"cartId"`
	Cart      Cart
	PaymentId int `db:"paymentId"`
	Payment   Payment
	CreatedAt time.Time `db:"createdAt"`
}

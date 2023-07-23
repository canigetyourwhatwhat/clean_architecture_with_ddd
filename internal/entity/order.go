package entity

import "time"

type Order struct {
	ID        int       `db:"id"`
	UserId    int       `db:"userId"`
	CartId    int       `db:"cartId"`
	PaymentId int       `db:"paymentId"`
	CreatedAt time.Time `db:"createdAt"`
}

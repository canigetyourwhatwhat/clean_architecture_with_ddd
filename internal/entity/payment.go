package entity

import (
	"time"
)

type Payment struct {
	ID        int           `db:"id"`
	UserId    int           `db:"userId"`
	Amount    float32       `db:"amount"`
	Method    PaymentMethod `db:"method"`
	CreatedAt time.Time     `db:"createdAt"`
	UpdatedAt time.Time     `db:"updatedAt"`
}

type PaymentMethod int

const (
	Card PaymentMethod = iota
	Cash
)

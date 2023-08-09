package entity

import (
	"time"
)

type Cart struct {
	ID         int        `db:"id"`
	UserId     int        `db:"userId"`
	TotalPrice float32    `db:"totalPrice"`
	TaxPrice   float32    `db:"taxPrice"`
	NetPrice   float32    `db:"netPrice"`
	Status     CartStatus `db:"status"`
	CartItems  []CartItem
	CreatedAt  time.Time `db:"createdAt"`
	UpdatedAt  time.Time `db:"updatedAt"`
}

type CartStatus int

const (
	InProgress CartStatus = iota
	Completed
)

package entity

import (
	"time"
)

type CartItem struct {
	ID          int    `db:"id"`
	ProductCode string `db:"productCode"`
	Product     *Product
	CartId      int       `db:"cartId"`
	Quantity    int       `db:"quantity"`
	TotalPrice  float32   `db:"totalPrice"`
	TaxPrice    float32   `db:"taxPrice"`
	NetPrice    float32   `db:"netPrice"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

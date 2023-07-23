package entity

import (
	"time"
)

type Product struct {
	ID               string    `db:"id"`
	Code             string    `db:"code"`
	Name             string    `db:"name"`
	Price            float32   `db:"price"`
	Stock            int       `db:"stock"`
	ShortDescription string    `db:"short_description"`
	LongDescription  string    `db:"long_description"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

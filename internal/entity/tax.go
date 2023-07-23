package entity

type Tax struct {
	ID   int     `db:"id"`
	Code string  `db:"code"`
	Rate float32 `db:"rate"`
}

const (
	DefaultTax = 0.2
)

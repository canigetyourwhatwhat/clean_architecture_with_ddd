package entity

import "time"

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"firstName" json:"FirstName"`
	LastName  string    `db:"lastName" json:"LastName"`
	Username  string    `db:"username" json:"Username"`
	Password  string    `db:"password" json:"Password"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

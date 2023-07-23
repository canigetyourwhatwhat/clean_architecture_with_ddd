package entity

import "time"

type Session struct {
	ID        string    `db:"id"`
	UserID    int       `db:"userId"`
	ExpiresAt time.Time `db:"expiresAt"`
}

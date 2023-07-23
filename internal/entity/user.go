package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"firstName" json:"FirstName"`
	LastName  string    `db:"lastName" json:"LastName"`
	Username  string    `db:"username" json:"Username"`
	Password  string    `db:"password" json:"Password"`
	TaxId     int       `db:"taxId"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (u *User) SetHashedPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

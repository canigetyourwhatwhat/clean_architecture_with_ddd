package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) {

	for query, errMsg := range map[string]string{
		createProductsTable: "Failed to create product table: %v",
		createTaxTable:      "Failed to create tax table: %v",
		createUserTable:     "Failed to create user table: %v",
		createSessionTable:  "Failed to create session table: %v",
		createCartTable:     "Failed to create cart table: %v",
		createCartItemTable: "Failed to create cart item table: %v",
		createPaymentTable:  "Failed to create payment table: %v",
		createOrderTable:    "Failed to create order table: %v",
	} {
		if _, err := db.Exec(query); err != nil {
			panic(fmt.Sprintf(errMsg, err.Error()))
		}
	}
}

func SeedTable(db *sqlx.DB) {
	var count int
	err := db.Get(&count, "select count(id) from Products")
	if err != nil {
		panic(fmt.Sprintf("Failed to get count of Products: %v", err))
	}

	if _, err = db.Exec("INSERT INTO tax (id, code, rate) SELECT 1, 'default', 0.2 WHERE NOT EXISTS (SELECT 1 FROM tax WHERE id = 1);"); err != nil {
		panic(fmt.Sprintf("Failed to insert default tax: %v", err))
	}

	if count < 30 {
		if _, err = db.Exec(insertProductsRecords); err != nil {
			panic(fmt.Sprintf("Failed to populate products record: %v", err))
		}
	}
}

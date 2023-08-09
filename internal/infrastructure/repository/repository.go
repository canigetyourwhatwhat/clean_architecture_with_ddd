package repository

import (
	"clean_architecture_with_ddd/internal/interface/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	DB *sqlx.DB
}

func NewRepo(db *sqlx.DB) repository.Repository {
	return Repo{DB: db}
}

func (r Repo) GetTaxRateById(id int) (tax float32, err error) {
	if err = r.DB.Get(&tax, "select rate from tax where id = ?", id); err != nil {
		return 0, err
	}
	return tax, nil
}

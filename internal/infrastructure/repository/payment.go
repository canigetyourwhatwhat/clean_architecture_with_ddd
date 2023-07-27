package repository

import "clean_architecture_with_ddd/internal/entity"

func (r Repo) GetPaymentById(id int) (p *entity.Payment, err error) {
	if err = r.DB.Get(p, "select * from payments where id = ?", id); err != nil {
		return nil, err
	}
	return p, nil
}

func (r Repo) CreatePayment(payment entity.Payment) (int, error) {
	query := `INSERT INTO payments (amount, userId, method) VALUES (:amount, :userId, :method)`
	result, err := r.DB.NamedExec(query, payment)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

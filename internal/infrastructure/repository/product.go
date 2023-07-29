package repository

import "clean_architecture_with_ddd/internal/entity"

func (r Repo) GetProductByCode(code string) (*entity.Product, error) {
	var p entity.Product
	if err := r.DB.Get(&p, "select * from products where code = ?", code); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r Repo) GetProductCount() (int, error) {
	var count int
	err := r.DB.Get(&count, "select count(id) from Products")
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (r Repo) ListProductsByPageNum(pageNum int, perPage int) ([]*entity.Product, error) {
	var products []*entity.Product
	offset := (pageNum - 1) * perPage

	err := r.DB.Select(&products, "select * from Products order by created_at desc limit ? offset ?", perPage, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

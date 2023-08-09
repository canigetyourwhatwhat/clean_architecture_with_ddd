package inputPort

import (
	"errors"
	"strconv"
)

func ProductCode(productCode string) error {
	if productCode == "" {
		return errors.New("product code is empty")
	}
	return nil
}

func ListProductsByPage(pageStr string) (page int, err error) {
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return -1, errors.New("page number is not integer")
	}
	if page <= 0 {
		return -1, errors.New("page number should not be negative")
	}
	return
}

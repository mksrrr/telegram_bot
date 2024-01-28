package product

import (
	"errors"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(index int64) (*Product, error) {
	if index >= int64(len(allProducts)) || index < 0 {
		indexStr := strconv.FormatInt(index, 10)
		errText := "Error. Can`t take \"" + indexStr + "\" element. There is no element with index \"" + indexStr + "\""
		return &Product{}, errors.New(errText)
	}

	return &allProducts[index], nil
}

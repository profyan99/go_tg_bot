package product

import (
	"errors"
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service *Service) List() []Product {
	return allProducts
}

func (service *Service) Get(id int) (*Product, error) {
	if id < 0 || id >= len(allProducts) {
		return nil, errors.New("invalid id")
	}

	log.Println("Get ID: ", id)

	return &allProducts[id], nil
}

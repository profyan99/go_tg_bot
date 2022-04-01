package product

import (
	"errors"
	"log"
	"math"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service *Service) List(offset uint, limit uint) ([]Product, error) {
	log.Printf("\nService list: %d %d", offset, limit)
	if err := service.validateIndex(offset); err != nil {
		return nil, err
	}

	endIndex := int(math.Min(float64(len(allProducts)), float64(offset+limit)))

	return allProducts[offset:endIndex], nil
}

func (service *Service) Count() int {
	return len(allProducts)
}

func (service *Service) Get(index uint) (*Product, error) {
	if err := service.validateIndex(index); err != nil {
		return nil, err
	}

	return &allProducts[index], nil
}

func (service *Service) Add(product *Product) uint {
	allProducts = append(allProducts, *product)

	return uint(len(allProducts) - 1)
}

func (service *Service) Remove(index uint) error {
	if err := service.validateIndex(index); err != nil {
		return err
	}

	allProducts = append(allProducts[:index], allProducts[index+1:]...)

	return nil
}

func (service *Service) Update(index uint, product *Product) error {
	if err := service.validateIndex(index); err != nil {
		return err
	}

	allProducts[index] = *product

	return nil
}

func (service *Service) validateIndex(index uint) error {
	if int(index) >= len(allProducts) {
		return errors.New("invalid id")
	}

	return nil
}

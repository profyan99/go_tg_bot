package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service *Service) List() []Product {
	return allProducts
}

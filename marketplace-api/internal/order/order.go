package order

import "Email-API/internal/products"

type OrderServiceDeps struct {
	Repo *OrderRepository
}

type OrderService struct {
	Repo *OrderRepository
}

func NewOrderService(deps *OrderServiceDeps) *OrderService {
	return &OrderService{
		Repo: deps.Repo,
	}
}

func (service *OrderService) Create(order *Order) (*Order, error) {
	result := service.Repo.Repo.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (service *OrderService) GetProducts(ids []uint) ([]products.Product, error) {
	var products []products.Product
	result := service.Repo.Repo.Find(&products, "id = ?", ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

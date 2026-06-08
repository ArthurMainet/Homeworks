package order

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
	order, err := service.Repo.Create(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

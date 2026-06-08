package order

import "Email-API/packages/db"

type OrderRepository struct {
	Repo *db.DB
}

func NewOrderRepository(db *db.DB) *OrderRepository {
	return &OrderRepository{
		Repo: db,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	//	var order Order
	result := repo.Repo.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetOrderById(id uint) (*Order, error) {
	var order Order
	result := repo.Repo.DB.Find(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepository) GetOrderByUserId(id uint) (*[]Order, error) {
	var order []Order
	result := repo.Repo.DB.Find(&order, "user_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

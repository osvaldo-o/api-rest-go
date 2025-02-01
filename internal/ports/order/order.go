package ports

import "flowers-mago/api/internal/domain"

type OrderService interface {
	Get() (orders []*domain.Order, err error)
	GetById(id *int) (*domain.Order, error)
	Create(order *domain.Order) (err error)
	Update(order *domain.Order) error
	Delete(id *string) error
}

type OrderRepository interface {
	GetAll() (orders []*domain.Order, err error)
	GetById(id *int) (*domain.Order, error)
	Insert(order *domain.Order) (err error)
	Update(order *domain.Order) error
	Delete(id *string) error
}

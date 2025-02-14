package ports

import "flowers-mago/api/internal/domain"

type OrderService interface {
	Get() (orders []*domain.Order, err error)
	GetById(id *int) (*domain.Order, error)
	Create(newOrder *domain.Order) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id *string) error
	SetStatus(status *bool, id *int) error
}

type OrderRepository interface {
	GetAll() (orders []*domain.Order, err error)
	GetById(id *int) (*domain.Order, error)
	Insert(order *domain.Order) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id *string) error
	SetStatus(status *bool, id *int) error
}

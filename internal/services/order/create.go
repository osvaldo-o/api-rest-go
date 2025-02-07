package order

import (
	"flowers-mago/api/internal/domain"
)

func (s *Service) Create(order *domain.Order) (*domain.Order, error) {
	return s.Repo.Insert(order)
}

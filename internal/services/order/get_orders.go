package order

import (
	"flowers-mago/api/internal/domain"
)

func (s *Service) Get() (orders []*domain.Order, err error) {
	_orders, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return _orders, nil
}

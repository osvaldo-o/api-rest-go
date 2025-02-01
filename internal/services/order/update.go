package order

import "flowers-mago/api/internal/domain"

func (s *Service) Update(order *domain.Order) error {
	return s.Repo.Update(order)
}

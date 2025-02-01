package order

import (
	"flowers-mago/api/internal/domain"
)

func (s *Service) GetById(id *int) (*domain.Order, error) {
	return s.Repo.GetById(id)
}

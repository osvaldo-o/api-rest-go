package order

import ports "flowers-mago/api/internal/ports/order"

var _ ports.OrderService = &Service{}

type Service struct {
	Repo ports.OrderRepository
}

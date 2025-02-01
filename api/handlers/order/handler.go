package order

import ports "flowers-mago/api/internal/ports/order"

type Handler struct {
	OrderService ports.OrderService
}

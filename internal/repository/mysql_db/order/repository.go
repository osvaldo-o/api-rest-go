package order

import (
	"database/sql"
	ports "flowers-mago/api/internal/ports/order"
)

var _ ports.OrderRepository = &Repository{}

type Repository struct {
	DB *sql.DB
}

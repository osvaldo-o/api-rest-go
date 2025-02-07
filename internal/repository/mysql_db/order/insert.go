package order

import "flowers-mago/api/internal/domain"

func (r *Repository) Insert(order *domain.Order) (*domain.Order, error) {

	query := `
		INSERT INTO orders (status, name, phone, delivery_date, delivery_time, place_delivery, description, price, comment)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id
	`

	var orderID int
	err := r.DB.QueryRow(
		query,
		order.Status,
		order.Name,
		order.Phone,
		order.DeliveryDate,
		order.DeliveryTime,
		order.PlaceDelivery,
		order.Description,
		order.Price,
		order.Comment,
	).Scan(&orderID)

	if err != nil {
		return nil, err
	}

	order.ID = orderID
	return order, nil

}

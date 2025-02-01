package order

import "flowers-mago/api/internal/domain"

func (r *Repository) GetById(id *int) (*domain.Order, error) {

	query := `
		SELECT id, status, name, phone, delivery_date, delivery_time, place_delivery, description, price, comment
		FROM orders
		WHERE id = ?
	`

	var order domain.Order
	err := r.DB.QueryRow(query, id).Scan(
		&order.ID,
		&order.Status,
		&order.Name,
		&order.Phone,
		&order.DeliveryDate,
		&order.DeliveryTime,
		&order.PlaceDelivery,
		&order.Description,
		&order.Price,
		&order.Comment,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil

}

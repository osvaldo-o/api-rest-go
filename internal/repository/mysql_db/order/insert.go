package order

import "flowers-mago/api/internal/domain"

func (r *Repository) Insert(order *domain.Order) error {

	query := `
		INSERT INTO orders (status, name, phone, delivery_date, delivery_time, place_delivery, description, price, comment)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
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
	)

	if err != nil {
		return err
	}

	return nil

}

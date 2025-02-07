package order

import (
	"flowers-mago/api/internal/domain"
	"log"
)

func (r *Repository) Insert(order *domain.Order) (*domain.Order, error) {

	query := `
		INSERT INTO orders (status, name, phone, delivery_date, delivery_time, place_delivery, description, price, comment)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
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
		log.Println(err)
		return nil, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	order.ID = int(orderID)

	return order, nil

}

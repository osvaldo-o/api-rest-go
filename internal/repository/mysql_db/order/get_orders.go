package order

import "flowers-mago/api/internal/domain"

func (r *Repository) GetAll() (orders []*domain.Order, err error) {

	rows, err := r.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsOrders []*domain.Order
	for rows.Next() {
		order := new(domain.Order)
		err := rows.Scan(
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
		newsOrders = append(newsOrders, order)
	}

	return newsOrders, nil

}

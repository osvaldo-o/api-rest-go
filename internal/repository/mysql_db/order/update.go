package order

import "flowers-mago/api/internal/domain"

func (r *Repository) Update(order *domain.Order) error {

	query := `UPDATE orders
			SET status = ?, name = ?, phone = ?, delivery_date = ?, delivery_time = ?, 
				place_delivery = ?, description = ?, price = ?, comment = ?
			WHERE id = ?`

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
		order.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}

	return nil

}

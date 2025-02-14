package order

func (r *Repository) SetStatus(status *bool, id *int) error {

	query := `UPDATE orders
	SET status = ?
	WHERE id = ?`

	result, err := r.DB.Exec(
		query,
		status,
		id,
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

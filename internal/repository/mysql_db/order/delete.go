package order

func (r *Repository) Delete(id *string) error {

	query := "DELETE FROM orders WHERE id = ?"

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}

	return nil

}

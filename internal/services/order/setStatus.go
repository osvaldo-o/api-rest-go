package order

func (s *Service) SetStatus(status *bool, id *int) error {
	return s.Repo.SetStatus(status, id)
}

package linkservice

func (s *Service) DeleteLink(id uint) (uint, error) {
	return s.lr.DeleteLink(id)
}

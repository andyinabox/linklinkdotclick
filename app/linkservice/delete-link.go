package linkservice

func (s *Service) DeleteLink(userId uint, id uint) (uint, error) {
	return s.lr.DeleteLink(userId, id)
}

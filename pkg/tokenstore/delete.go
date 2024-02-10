package tokenstore

func (s *Store) Delete(hash string) error {
	tx := s.db.Where("hash = ?", hash).Delete(&Token{})
	return tx.Error
}

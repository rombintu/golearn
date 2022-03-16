package store

func (s *Store) CreateDeclaration(d Declaration) error {
	return s.Database.Create(&d).Error
}

func (s *Store) DeleteDeclaration(userID int, title string) error {
	if title != "" {
		return s.Database.Delete(&Declaration{}, "user_id = ? AND title = ?", userID, title).Error
	}
	return s.Database.Delete(&Declaration{}, "user_id = ?", userID).Error
}

func (s *Store) GetDeclarationByUserID(userID int) ([]Declaration, error) {
	var d []Declaration
	if err := s.Database.Find(&d, "user_id = ?", userID).Error; err != nil {
		return []Declaration{}, err
	}
	return d, nil
}

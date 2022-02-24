package store

func (s *Store) CreateStudentGroup(group StudentGroup) error {
	return s.Database.Create(&group).Error
}

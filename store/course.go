package store

func (s *Store) CreateCourse(crs Course) (Course, error) {
	tx := s.Database.Create(&crs)
	return crs, tx.Error
}

func (s *Store) DeleteCourse(crsId int) error {
	return s.Database.Delete(&Course{}, "ID = ?", crsId).Error
}

func (s *Store) GetAllCourses() ([]Course, error) {
	var crsAll []Course
	if err := s.Database.Find(&crsAll).Error; err != nil {
		return []Course{}, err
	}
	return crsAll, nil
}

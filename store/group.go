package store

func (s *Store) CreateGroup(group Group) error {
	return s.Database.Create(&group).Error
}

func (s *Store) DeleteGroupByTitle(groupTitle string) error {
	return s.Database.Delete(&Group{}, "title = ?", groupTitle).Error
}

func (s *Store) DeleteGroupByID(groupID int) error {
	return s.Database.Delete(&Group{}, groupID).Error
}

func (s *Store) JoinUserGroup(userLogin, groupName string) error {
	var user User
	var group Group
	if err := s.Database.First(&user, "login = ?", userLogin).Error; err != nil {
		return err
	}

	if err := s.Database.First(&group, "title = ?", groupName).Error; err != nil {
		return err
	}

	groupUsers := GroupUsers{
		UserID:  int(user.ID),
		GroupID: int(group.ID),
	}

	return s.Database.Create(&groupUsers).Error
}

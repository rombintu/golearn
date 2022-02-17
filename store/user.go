package store

func (s *Store) CreateUser(user User) error {
	return s.Database.Create(&User{
		Account:  user.Account,
		Password: user.Password,
		Role:     user.Role,
	}).Error
}

func (s *Store) GetUser(id int) (User, error) {
	var user User
	err := s.Database.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// TODO Unique account
func (s *Store) GetUserByAccount(account string) (User, error) {
	var user User
	err := s.Database.First(&user, "account = ?", account).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

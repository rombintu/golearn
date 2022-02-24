package store

import "github.com/rombintu/golearn/tools"

func (s *Store) CreateUser(user User) error {
	hashPass, err := tools.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	return s.Database.Create(&user).Error
}

func (s *Store) GetUserByID(id int) (User, error) {
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

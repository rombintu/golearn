package store

import (
	"github.com/rombintu/golearn/tools"
)

func (s *Store) CreateUser(user User) (User, error) {
	hashPass, err := tools.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = hashPass
	tx := s.Database.Create(&user)
	return user, tx.Error
}

func (s *Store) UpdateUser(user User) (User, error) {
	var userOld User
	if err := s.Database.First(&userOld, "id = ?", user.ID).Error; err != nil {
		return User{}, err
	}

	// const layoutDate = "01.01.2001"
	// bthDay, err := time.Parse(layoutDate, user.DateOfBirth)

	userOld.FirstName = user.FirstName
	userOld.LastName = user.LastName
	userOld.Address = user.Address
	userOld.Mail = user.Mail
	userOld.Phone = user.Phone
	userOld.DateOfBirth = user.DateOfBirth

	if err := s.Database.Save(&userOld).Error; err != nil {
		return User{}, err
	}

	return userOld, nil
}

func (s *Store) DeleteUser(idUser int) error {
	var user User
	return s.Database.Delete(&user, "id = ?", idUser).Error
}

func (s *Store) CreateWorker(user Worker) error {
	hashPass, err := tools.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	return s.Database.Create(&user).Error
}

func (s *Store) CreateTeacher(user Teacher) error {
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

func (s *Store) GetWorkerByID(id int) (Worker, error) {
	var user Worker
	err := s.Database.First(&user, id).Error
	if err != nil {
		return Worker{}, err
	}
	return user, nil
}

func (s *Store) GetTeacherByID(id int) (Teacher, error) {
	var user Teacher
	err := s.Database.First(&user, id).Error
	if err != nil {
		return Teacher{}, err
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

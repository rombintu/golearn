package store

import (
	"fmt"
	"log"

	"github.com/rombintu/golearn/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

type Store struct {
	Database *gorm.DB
	Config   *config.Postgres
}

// Struct for test
type Ping struct {
	Message string `json:"message"`
}

func (s *Store) Open() error {
	if !s.Config.Dev {
		connStr := fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			s.Config.User,
			s.Config.Password,
			s.Config.Dbname,
			s.Config.SSLMode,
		)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			return err
		}
		s.Database = db
		return nil
	}
	db, err := gorm.Open(sqlite.Open("./store/dev.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	s.Database = db
	return nil
}

func (s *Store) Close() {
	db, err := s.Database.DB()
	if err != nil {
		log.Println(err)
	}
	db.Close()
}

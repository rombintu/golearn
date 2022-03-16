package store

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Content string
}

// User model for db
type User struct {
	gorm.Model
	Account     string        `json:"account" gorm:"unique"`
	Password    string        `json:"password"`
	Role        string        `json:"role"`
	Address     string        `json:"address"`
	Mail        string        `json:"mail"`
	Phone       string        `json:"phone"`
	DateOfBirth time.Duration `json:"date_of_birth"`
	Declaration []Declaration `gorm:"foreignKey:UserID"`
}

type Worker struct {
	User
	Experience int    `json:"experience"`
	Position   string `json:"position"`
}

type Teacher struct {
	Worker
	Education     string `json:"education"`
	AcademicTitle string `json:"academic_title"`
	AcademicSte   string `json:"academic_ste"`
}

type OfficePaper struct {
	gorm.Model
	Title string `json:"title"`
}

type Declaration struct {
	OfficePaper
	Date    time.Duration `json:"date"`
	Content string        `json:"content"`
	UserID  int           `json:"user_id"`
}

type Course struct {
	OfficePaper
}

type Journal struct {
	OfficePaper
	StudentPresence bool `json:"student_presence"`
	Assessment      int  `json:"assessment"`
}

type Contract struct {
	OfficePaper
	Date           time.Duration `json:"date"`
	ValidatyPeriod time.Duration `json:"validaty_period"`
	Content        string        `json:"content"`
	Amount         int           `json:"amount"`
	UserRefer      int
	User           User `gorm:"foreignKey:UserRefer"`
	WorkerRefer    int
	Worker         Worker `gorm:"foreignKey:WorkerRefer"`
}

type Payment struct {
	OfficePaper
	Amount int           `json:"amount"`
	Date   time.Duration `json:"date"`
}

type Services struct {
	OfficePaper
	DateStart time.Duration `json:"date_start"`
	DateEnd   time.Duration `json:"date_end"`
}

type Refferal struct {
	OfficePaper
	ServiceEnd    time.Duration `json:"service_end"`
	ServiceAmount int           `json:"service_amount"`
}

type Plan struct {
	OfficePaper
	DateStart    time.Duration
	DateEnd      time.Duration
	Kind         string
	Place        string
	Agree        bool
	DateTransfer time.Duration
	GroupRefer   int
	Group        Group `gorm:"foreignKey:GroupRefer"`
	TeacherRefer int
	Teacher      Teacher `gorm:"foreignKey:TeacherRefer"`
	CourseRefer  int
	Course       Course `gorm:"foreignKey:CourceRefer"`
}

type Group struct {
	gorm.Model
	Title            string        `json:"title" gorm:"unique"`
	DateOfEnrollment time.Duration `json:"date_of_enrollment"`
	DateOfDeducation time.Duration `json:"date_od_deducation"`
	Users            []User        `gorm:"many2many:group_users;"`
}

type GroupUsers struct {
	UserID  int `gorm:"primaryKey"`
	GroupID int `gorm:"primaryKey"`
}

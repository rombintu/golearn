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
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Address     string        `json:"address"`
	Mail        string        `json:"mail"`
	Phone       string        `json:"phone"`
	DateOfBirth time.Time     `json:"date_of_birth"`
	Declaration []Declaration `gorm:"foreignKey:UserID"`
	Token       string        `gorm:"migration" json:"token"`
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
	Title string `json:"title" gorm:"unique"`
}

type Declaration struct {
	OfficePaper
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
	UserID  int       `json:"user_id"`
}

type Course struct {
	OfficePaper
	Abount   string `json:"about"`
	Tags     string `json:"tags"`
	IsActive bool   `json:"is_active"`
}

type Journal struct {
	OfficePaper
	StudentPresence bool `json:"student_presence"`
	Assessment      int  `json:"assessment"`
}

type Contract struct {
	OfficePaper
	Date           time.Time `json:"date"`
	ValidatyPeriod time.Time `json:"validaty_period"`
	Content        string    `json:"content"`
	Amount         int       `json:"amount"`
	UserRefer      int
	User           User `gorm:"foreignKey:UserRefer"`
	WorkerRefer    int
	Worker         Worker `gorm:"foreignKey:WorkerRefer"`
}

type Payment struct {
	OfficePaper
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

type Services struct {
	OfficePaper
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}

type Refferal struct {
	OfficePaper
	ServiceEnd    time.Time `json:"service_end"`
	ServiceAmount int       `json:"service_amount"`
}

type Plan struct {
	OfficePaper
	DateStart    time.Time
	DateEnd      time.Time
	Kind         string
	Place        string
	Agree        bool
	DateTransfer time.Time
	GroupRefer   int
	Group        Group `gorm:"foreignKey:GroupRefer"`
	TeacherRefer int
	Teacher      Teacher `gorm:"foreignKey:TeacherRefer"`
	CourseRefer  int
	Course       Course `gorm:"foreignKey:CourceRefer"`
}

type Group struct {
	gorm.Model
	Title            string    `json:"title" gorm:"unique"`
	DateOfEnrollment time.Time `json:"date_of_enrollment"`
	DateOfDeducation time.Time `json:"date_od_deducation"`
	Users            []User    `gorm:"many2many:group_users;"`
}

type GroupUsers struct {
	UserID  int `gorm:"primaryKey"`
	GroupID int `gorm:"primaryKey"`
}

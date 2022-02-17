package store

import (
	"time"

	"gorm.io/gorm"
)

// User model for db
type User struct {
	// Id          int           `json:"id"`
	gorm.Model
	Account     string        `json:"account"`
	Password    string        `json:"password"`
	Role        string        `json:"role"`
	Address     string        `json:"address"`
	Mail        string        `json:"mail"`
	Phone       string        `json:"phone"`
	DateOfBirth time.Duration `json:"date_of_birth"`
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
	// Id    int    `json:"id"`
	gorm.Model
	Title string `json:"title"`
}

type Declaration struct {
	OfficePaper
	Date    time.Duration `json:"date"`
	Content string        `json:"content"`
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

type Group struct {
	// Id    int    `json:"id"`
	gorm.Model
	Title string `json:"title"`
}

type StudentGroup struct {
	Group
	DateOfEnrollment time.Duration `json:"date_of_enrollment"`
	DateOfDeducation time.Duration `json:"date_od_deducation"`
}

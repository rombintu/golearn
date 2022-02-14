package store

// User model for db
type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

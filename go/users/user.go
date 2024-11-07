package users

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"is_active"`
}

package dto

import "time"

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Age      uint   `json:"age"`
}

type UserResponse struct {
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Age       uint      `json:"age"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package dto

import "time"

type RegisterRequest struct {
	Email    string
	Password string
	Fullname string
	Age      uint
}

type RegisterResponse struct {
	Email     string
	Fullname  string
	Age       uint
	Role      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginRequest struct {
	Email    string
	Password string
}

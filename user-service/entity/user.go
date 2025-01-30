package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `gorm:"unique"`
	Password string
	Fullname string
	Age      uint
	Role     string
	Active   bool
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Role = "customer"
	u.Active = true

	return u.Base.BeforeCreate(tx)
}

func (u *User) HashPassword() error {
	if u.Password == "" {
		return errors.New("cant hash password: empty string")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)

	return nil
}

func (u *User) ComparePassword(inputPassword string) error {
	byteInputPassword := []byte(inputPassword)
	hashCurrentPassword := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(hashCurrentPassword, byteInputPassword)
	if err != nil {
		return err
	}

	return nil
}

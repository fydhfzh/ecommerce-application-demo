package repository

import (
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	SaveUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	result := u.db.First(&user, "email = ?", email)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) SaveUser(user entity.User) (*entity.User, error) {
	result := u.db.Save(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

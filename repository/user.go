package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	StoreProfile(userProfile entity.UserProfile) (entity.UserProfile, error)
	Store(user entity.User) (entity.User, error)
	Login(email string) (entity.User, error)
	StoreAddress(userAddress entity.UserAddress) (entity.UserAddress, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) StoreProfile(userProfile entity.UserProfile) (entity.UserProfile, error) {
	if err := r.db.Debug().Create(&userProfile).Error; err != nil {
		return userProfile, err
	}

	return userProfile, nil
}

func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepository) Login(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepository) StoreAddress(userAddress entity.UserAddress) (entity.UserAddress, error) {
	if err := r.db.Debug().Create(&userAddress).Error; err != nil {
		return userAddress, err
	}

	return userAddress, nil
}

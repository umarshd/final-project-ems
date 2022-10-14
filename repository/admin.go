package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	Store(admin entity.Admin) (entity.Admin, error)
	Login(email string) (entity.Admin, error)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db}
}

func (r AdminRepository) Store(admin entity.Admin) (entity.Admin, error) {
	if err := r.db.Debug().Create(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (r AdminRepository) Login(email string) (entity.Admin, error) {
	var admin entity.Admin
	if err := r.db.Debug().Where("email = ?", email).First(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

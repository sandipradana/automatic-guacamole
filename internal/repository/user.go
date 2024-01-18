package repository

import (
	"automatic-guacamole/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(db *gorm.DB) ([]model.User, error)
	GetByID(db *gorm.DB, id uint64) (*model.User, error)
	Create(db *gorm.DB, User *model.User) error
	Update(db *gorm.DB, User *model.User) error
	Delete(db *gorm.DB, id uint64) error
	GetUserByEmail(db *gorm.DB, email string) (*model.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll(db *gorm.DB) ([]model.User, error) {
	var items []model.User
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *userRepository) GetByID(db *gorm.DB, id uint64) (*model.User, error) {
	var item model.User
	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *userRepository) Create(db *gorm.DB, User *model.User) error {
	if err := db.Create(User).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(db *gorm.DB, User *model.User) error {
	if err := db.Save(User).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(db *gorm.DB, id uint64) error {
	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	var User model.User
	if err := db.Where("email = ?", email).First(&User).Error; err != nil {
		return nil, err
	}
	return &User, nil
}

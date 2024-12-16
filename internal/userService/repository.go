package userService

import (
	"firstProject/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUserById(id uint, updatedUser models.User) (models.User, error)
	DeleteUserById(id uint) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) UpdateUserById(id uint, updatedUser models.User) (models.User, error) {
	findByID := r.db.First(&models.User{}, id)
	if findByID.Error != nil {
		return updatedUser, findByID.Error
	}
	updatedUser.ID = id
	result := r.db.Model(&updatedUser).Update("email", updatedUser.Email)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return updatedUser, nil
}
func (r *userRepository) DeleteUserById(id uint) error {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&models.User{}, id).Error
}

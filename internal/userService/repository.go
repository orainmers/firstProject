package userService

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

type UserRepository interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUserById(id uint, updatedUser User) (User, error)
	DeleteUserById(id uint) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) UpdateUserById(id uint, updatedUser User) (User, error) {
	findByID := r.db.First(&User{}, id)
	if findByID.Error != nil {
		return updatedUser, findByID.Error
	}
	updatedUser.ID = id
	result := r.db.Model(&updatedUser).Update("email", updatedUser.Email)
	if result.Error != nil {
		return User{}, result.Error
	}
	return updatedUser, nil
}
func (r *userRepository) DeleteUserById(id uint) error {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&User{}, id).Error
}

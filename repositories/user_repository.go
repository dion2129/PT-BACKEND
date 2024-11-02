package repositories

import (
	"api-test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(id int) (models.User, error)
	FindAll() ([]models.User, error)
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}

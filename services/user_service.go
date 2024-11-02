package services

import (
	"api-test/models"
	"api-test/repositories"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.userRepo.Save(user)
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *userService) GetUserByID(id int) (models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}

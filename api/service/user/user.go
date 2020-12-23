package service

import (
	"errors"
	"go-practice/api/repository"
	repository "go-practice/api/repository/user"
	"go-practice/models"
)

type UserService interface {
	AddUser(models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(int) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(int) (models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (u *userService) AddUser(user models.User) (models.User, error) {
	_, err := u.userRepository.GetUserByID(user.ID)
	if err == nil {
		return user, errors.New("this is login")
	}
	return u.userRepository.AddUser(user)
}
func (u *userService) GetAllUsers() (users []models.User, err error) {
	return u.userRepository.GetAllUsers()
}
func (u *userService) GetUserByID(id string) (models.User, error) {
	return u.userRepository.GetUserByID(id)
}
func (u *userService) GetUserByEmail(email string) (models.User, error) {
	return u.userRepository.GetUserByEmail(email)
}
func (u *userService) UpdateUser(user models.User) (models.User, error) {
	return u.userRepository.UpdateUser(user)

}
func (u *userService) DeleteUser(id int) (models.User, error) {
	return u.userRepository.DeleteUser(id)
}

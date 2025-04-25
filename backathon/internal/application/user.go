package application

import "backathon/internal/domain"

type UserUsecaseInterface interface {
	GetUser(id int32) (domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(id int32, user domain.User) (domain.User, error)
	DeleteUser(id int32) error
}

type UserRepositoryInterface interface {
	GetUser(id int32) (domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(id int32, user domain.User) (domain.User, error)
	DeleteUser(id int32) error
}

type UserUsecase struct {
	userRepository UserRepositoryInterface
}

func NewUserUsecase(userRepository UserRepositoryInterface) UserUsecaseInterface {
	return &UserUsecase{
		userRepository: userRepository,
	}
}
func (u *UserUsecase) GetUser(id int32) (domain.User, error) {
	return u.userRepository.GetUser(id)
}

func (u *UserUsecase) CreateUser(user domain.User) (domain.User, error) {
	return u.userRepository.CreateUser(user)
}

func (u *UserUsecase) UpdateUser(id int32, user domain.User) (domain.User, error) {
	return u.userRepository.UpdateUser(id, user)
}

func (u *UserUsecase) DeleteUser(id int32) error {
	return u.userRepository.DeleteUser(id)
}

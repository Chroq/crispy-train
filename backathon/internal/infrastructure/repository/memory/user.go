package memory

import (
	"backathon/domain"
	"sync"
)

type UserRepository struct {
	storage sync.Map
	nextId  int32
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
func (r *UserRepository) GetUser(id int32) (domain.User, error) {
	if value, ok := r.storage.Load(id); ok {
		return value.(domain.User), nil
	}
	return domain.User{}, nil
}

func (r *UserRepository) CreateUser(user domain.User) (domain.User, error) {
	r.nextId++
	user.Id = r.nextId
	r.storage.Store(r.nextId, user)
	return user, nil
}

func (r *UserRepository) UpdateUser(id int32, user domain.User) (domain.User, error) {
	if _, ok := r.storage.Load(id); ok {
		r.storage.Store(id, user)
		return user, nil
	}
	return domain.User{}, nil
}
func (r *UserRepository) DeleteUser(id int32) error {
	if _, ok := r.storage.Load(id); ok {
		r.storage.Delete(id)
		return nil
	}
	return nil
}

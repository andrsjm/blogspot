package flow

import (
	"blogspot/entity"
	"blogspot/repository"
)

type userFlow struct {
	repo repository.IUserRepository
}

func NewUserFlow(userRepo repository.IUserRepository) *userFlow {
	return &userFlow{
		repo: userRepo,
	}
}

func (f *userFlow) Login(user entity.User) (entity.User, error) {
	user, err := f.repo.Login(user)
	return user, err
}

func (f *userFlow) Register(user entity.User) (err error) {
	err = f.repo.Register(user)
	return
}

func (f *userFlow) Update(user entity.User) (err error) {
	err = f.repo.Update(user)
	return
}

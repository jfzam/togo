package application

import (
	"github.com/jfzam/togo/domain/entity"
	"github.com/jfzam/togo/domain/repository"
)

type userApp struct {
	ur repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByUsernameAndPassword(*entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.ur.SaveUser(user)
}

func (u *userApp) GetUser(userId uint64) (*entity.User, error) {
	return u.ur.GetUser(userId)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.ur.GetUsers()
}

func (u *userApp) GetUserByUsernameAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.ur.GetUserByUsernameAndPassword(user)
}

package user

import (
	"github.com/thtrangphu/bookStoreSvc/entity"
)

type User interface {
	Create(e *entity.User) (entity.ID, error)
	Get(id entity.ID) (*entity.User, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	User
}

// UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(email, password, firstName, lastName string) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}

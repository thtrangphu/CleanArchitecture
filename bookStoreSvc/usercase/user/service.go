package user

import (
	"github.com/thtrangphu/bookStoreSvc/entity"
	"time"
)

type Service struct {
	repo Repository
}

// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateNewUser Create an user
func (s *Service) CreateNewUser(email, password, fullName string) (entity.ID, error) {
	e, err := entity.CreateNewUser(email, password, fullName)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetUser Get an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

// UpdateUser Update an user
func (s *Service) UpdateUser(e *entity.User) error {
	if e.CheckUserInfor() {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

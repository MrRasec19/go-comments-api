package service

import (
	"go-comments-api/internal/model"
	"go-comments-api/internal/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{store: s}
}

func (s *Service) GetAllUsers() ([]*model.User, error) {
	return s.store.GetAll()
}

func (s *Service) GetBooksByID(id int) (*model.User, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateUser(user *model.User) (*model.User, error) {
	return s.store.Create(user)
}

func (s *Service) Updateuser(id int, user *model.User) (*model.User, error) {
	return s.store.Update(id, user)
}

func (s *Service) Deleteuser(id int) error {
	return s.store.Delete(id)
}

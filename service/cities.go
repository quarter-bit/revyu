package service

import (
	"revyu/entities"
)

type Repo interface {
	Create(id int, name string, state string) error
	Delete(id int) error
	Update(id int, name string, state string) error
	List() ([]entities.City, error)
}

type service struct {
	repo Repo
}

func New(repo Repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(id int, name string, state string) error {
	return s.repo.Create(id, name, state)
}

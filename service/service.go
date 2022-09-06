package service

import "wire_best_practice/repo"

type Xlogic interface {
	GetEntity(id int) (*repo.Entity, error)
}

type Service struct {
	repo repo.Database
}

func New(r repo.Database) *Service {
	return &Service{repo: r}
}

func (s *Service) GetEntity(id int) (*repo.Entity, error) {
	entity, err := s.repo.GetEntity(id)
	if err != nil {
		return nil, err
	}

	// do something with entity...

	return entity, nil
}

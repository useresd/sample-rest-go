package main

import (
	"log"
)

type Service struct {
	log *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) *Service {
	return &Service{
		log: log,
		repo: repo,
	}
}

func (s *Service) GetAll() ([]*Currency, error) {
	s.log.Printf("info: finding all currencies\n")
	return s.repo.GetAll(), nil
}

func (s *Service) FindByID(id int) (*Currency, error) {
	s.log.Printf("info: finding currency by id %d\n", id)
	return s.repo.FindByID(id)
}

func (s *Service) Store(data struct{Name string; Abbr string}) (*Currency, error) {
	s.log.Printf("info: storing currency")
	currency := &Currency{
		Name: data.Name,
		Abbr: data.Abbr,
	}
	return s.repo.Save(currency)
}
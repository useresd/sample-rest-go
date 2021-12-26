package currency

import (
	"log"
)

type Service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

// Get all currencies.
func (s *Service) GetAll() ([]*Currency, error) {
	s.log.Printf("info: finding all currencies\n")
	return s.repo.GetAll(), nil
}

// FindByID finds a currency by its id
func (s *Service) FindByID(id int) (*Currency, error) {
	s.log.Printf("info: finding currency by id %d\n", id)
	return s.repo.FindByID(id)
}

// Store saves the currency into the datastore.
func (s *Service) Store(data EditDTO) (*Currency, error) {
	s.log.Printf("info: storing currency")

	// validate the data.
	if err := data.Validate(); err != nil {
		return nil, err
	}

	// construct the currency
	currency := &Currency{
		Name: data.Name,
		Abbr: data.Abbr,
	}
	return s.repo.Save(currency)
}

func (s *Service) Update(id int, data EditDTO) (*Currency, error) {
	s.log.Printf("info: updating currency")

	// find the old currency
	currency, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	currency.Name = data.Name
	currency.Abbr = data.Abbr

	modified, err := s.repo.Save(currency)
	if err != nil {
		return nil, err
	}

	return modified, nil
}

func (s *Service) Delete(id int) error {
	s.log.Printf("info: deleting currency")

	// find the old currency
	currency, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// remove the currency
	if err := s.repo.Delete(currency); err != nil {
		return err
	}

	return nil
}

package account

import (
	"fmt"
	"log"

	"github.com/useresd/ledger/app/currency"
)

type Service struct {
	log          *log.Logger
	repo         Repository
	currencyRepo currency.Repository
}

func NewService(log *log.Logger, repo Repository, currencyRepo currency.Repository) *Service {
	return &Service{
		log:          log,
		repo:         repo,
		currencyRepo: currencyRepo,
	}
}

func (s *Service) FindByID(id int) (*Account, error) {
	s.log.Printf("find account by id%d\n", id)
	return s.repo.FindByID(id)
}

func (s *Service) GetAll() ([]*Account, error) {
	s.log.Printf("getting all accounts")
	return s.repo.GetAll()
}

func (s *Service) Store(data EditDTO) (*Account, error) {
	s.log.Printf("saving new account")

	// Validate the data.
	if err := data.Validate(); err != nil {
		return nil, err
	}

	// Get the currency
	currency, err := s.currencyRepo.FindByID(data.CurrencyID)
	if err != nil {
		return nil, fmt.Errorf("could not find currency by id: %d", data.CurrencyID)
	}

	// Constructing the account
	account := &Account{
		Name:     data.Name,
		Currency: currency,
	}

	// Saving the account
	return s.repo.Save(account)
}

// Update updates the account data.
func (s *Service) Update(id int, data EditDTO) (*Account, error) {
	s.log.Printf("info: updating account")

	if err := data.Validate(); err != nil {
		return nil, err
	}

	// Get the account
	account, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find account with id: %d", id)
	}

	// Get the currency
	currency, err := s.currencyRepo.FindByID(data.CurrencyID)
	if err != nil {
		return nil, fmt.Errorf("could not find currency by id: %d", data.CurrencyID)
	}

	account.Name = data.Name
	account.Currency = currency

	modified, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return modified, nil
}

func (s *Service) Delete(id int) error {
	s.log.Printf("deleting account with id: %d", id)

	// finding account to delete
	account, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// delete the account and return
	return s.repo.Delete(account)
}

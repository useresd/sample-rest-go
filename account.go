package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Account entity
type Account struct {
	ID int `json:"account"`
	Name string `json:"name"`
	Currency *Currency `json:"currency"`
}

// AccountRepository Interface
type AccountRepository interface {
	FindByID(int) (*Account, error)
	GetAll() ([]*Account, error)
}

// Mocking The Repository Interface
type AccountMock struct {
	Data []*Account
}

func NewAccountMock() *AccountMock {
	return &AccountMock{}
}

func (r *AccountMock) Populate() *AccountMock {
	r.Data = append(r.Data, 
		&Account{ID: 1, Name: "Cash in Hand", Currency: &Currency{ID: 1, Abbr: "SDG"}},
		&Account{ID: 2, Name: "Cash in Bank", Currency: &Currency{ID: 2, Abbr: "USD"}},
		&Account{ID: 3, Name: "Receivables", Currency: &Currency{ID: 3, Abbr: "AED"}},
	)
	return r
}

func (r *AccountMock) FindByID(id int) (*Account, error) {

	for _, d := range r.Data {
		if d.ID == id {
			return d, nil
		}
	}

	return nil, fmt.Errorf("could not find account by id %d", id)
}

func (r *AccountMock) GetAll() ([]*Account, error) {
	return r.Data, nil
}

// -------------------------------------
//
// The Account Service
//
// -------------------------------------

type AccountService struct {
	log *log.Logger
	repo AccountRepository
}

func NewAccountService(log *log.Logger, repo AccountRepository) *AccountService {
	return &AccountService{
		log: log,
		repo: repo,
	}
}

func (s *AccountService) FindByID(id int) (*Account, error) {
	s.log.Printf("find account by id%d\n", id)
	return s.repo.FindByID(id)
}

func (s *AccountService) GetAll() ([]*Account, error) {
	s.log.Printf("getting all accounts")
	return s.repo.GetAll()
}


// -------------------------------------
//
// The Account APIs
//
// -------------------------------------

type AccountAPI struct{
	Service *AccountService
}

func NewAccountAPI(service *AccountService) *AccountAPI {
	return &AccountAPI{Service: service}
}

func (a *AccountAPI) FindByID(rw http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return err
	}
	account, err := a.Service.FindByID(id)
	if err != nil {
		return err
	}

	WriteJson(rw, http.StatusOK, account)
	return nil
}

func (a *AccountAPI) GetAll(rw http.ResponseWriter, r *http.Request) error {
	accounts, err := a.Service.GetAll()
	if err != nil {
		return err
	}

	WriteJson(rw, http.StatusOK, accounts)
	return nil
}
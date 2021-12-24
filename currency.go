package main

import (
	"fmt"
	"math/rand"
)

type Currency struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}

type Repository interface {
	FindByID(id int) (*Currency, error)
	GetAll() []*Currency
	Save(*Currency) (*Currency, error)
}

type RepoMock struct {
	Currencies []*Currency
}

func NewMock() *RepoMock {
	return &RepoMock{}
}

func (r *RepoMock) Populate() *RepoMock {
	r.Currencies = append(r.Currencies, 
		&Currency{ID: 1, Name: "Sudanese Pound", Abbr: "SDG"},
		&Currency{ID: 2, Name: "United States Dollars", Abbr: "USD"},
		&Currency{ID: 3, Name: "Egyption Pound", Abbr: "EGP"},
	)
	return r
}

func (r *RepoMock) FindByID(id int) (*Currency, error) {
	for _, d := range r.Currencies {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, fmt.Errorf("Currency with id %d could not be found", id)
}

func (r *RepoMock) GetAll() []*Currency {
	return r.Currencies
}

func (r *RepoMock) Save(currency *Currency) (*Currency, error) {

	// New Currency
	if currency.ID == 0 {
		id := rand.Int()
		currency.ID = id
		r.Currencies = append(r.Currencies, currency)
	}

	// Update currency
	for i, d := range r.Currencies {
		if d.ID == currency.ID {
			r.Currencies[i] = currency
		}
	}

	return currency, nil	
}
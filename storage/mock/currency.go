package mock

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/useresd/ledger/app/currency"
)

type Currency struct {
	Currencies []*currency.Currency
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (r *Currency) Populate() *Currency {
	r.Currencies = append(r.Currencies,
		&currency.Currency{ID: 1, Name: "Sudanese Pound", Abbr: "SDG"},
		&currency.Currency{ID: 2, Name: "United States Dollars", Abbr: "USD"},
		&currency.Currency{ID: 3, Name: "Egyption Pound", Abbr: "EGP"},
	)
	return r
}

func (r *Currency) FindByID(id int) (*currency.Currency, error) {
	for _, d := range r.Currencies {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, fmt.Errorf("currency with id %d could not be found", id)
}

func (r *Currency) GetAll() []*currency.Currency {
	return r.Currencies
}

func (r *Currency) Save(currency *currency.Currency) (*currency.Currency, error) {

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

func (r *Currency) Delete(c *currency.Currency) error {

	// Check of the currency id exists
	if c.ID == 0 {
		return errors.New("could not delete a currency that doesn't have an id")
	}

	// loop for currencies and delete the found one.
	for i, d := range r.Currencies {
		if d.ID == c.ID {
			r.Currencies = append(make([]*currency.Currency, 0), append(r.Currencies[:i], r.Currencies[i+1:]...)...)
			return nil
		}
	}

	return nil
}

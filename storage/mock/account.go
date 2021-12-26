package mock

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/useresd/ledger/app/account"
	"github.com/useresd/ledger/app/currency"
)

// Mocking The Repository Interface
type Account struct {
	Data []*account.Account
}

func NewAccount() *Account {
	return &Account{}
}

func (r *Account) Populate() *Account {
	r.Data = append(r.Data,
		&account.Account{ID: 1, Name: "Cash in Hand", Currency: &currency.Currency{ID: 1, Abbr: "SDG"}},
		&account.Account{ID: 2, Name: "Cash in Bank", Currency: &currency.Currency{ID: 2, Abbr: "USD"}},
		&account.Account{ID: 3, Name: "Receivables", Currency: &currency.Currency{ID: 3, Abbr: "AED"}},
	)
	return r
}

func (r *Account) FindByID(id int) (*account.Account, error) {

	for _, d := range r.Data {
		if d.ID == id {
			return d, nil
		}
	}

	return nil, fmt.Errorf("could not find account by id %d", id)
}

func (r *Account) GetAll() ([]*account.Account, error) {
	return r.Data, nil
}

func (r *Account) Save(account *account.Account) (*account.Account, error) {
	// New Currency
	if account.ID == 0 {
		id := rand.Int()
		account.ID = id
		r.Data = append(r.Data, account)
	}

	// Update currency
	for i, d := range r.Data {
		if d.ID == account.ID {
			r.Data[i] = account
		}
	}

	// return  the account
	return account, nil
}

func (r *Account) Delete(a *account.Account) error {

	// Check of the currency id exists
	if a.ID == 0 {
		return errors.New("could not delete a account that doesn't have an id")
	}

	// loop for currencies and delete the found one.
	for i, d := range r.Data {
		if d.ID == a.ID {
			r.Data = append(make([]*account.Account, 0), append(r.Data[:i], r.Data[i+1:]...)...)
			return nil
		}
	}

	return nil
}

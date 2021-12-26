package account

import (
	"errors"

	"github.com/useresd/ledger/app/currency"
)

// Account entity
type Account struct {
	ID       int                `json:"account"`
	Name     string             `json:"name"`
	Currency *currency.Currency `json:"currency"`
}

type EditDTO struct {
	Name       string `json:"name"`
	CurrencyID int    `json:"currency_id"`
}

func (e *EditDTO) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}

	if e.CurrencyID == 0 {
		return errors.New("currency_id is required")
	}

	return nil
}

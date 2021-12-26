package currency_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/useresd/ledger/app/currency"
	"github.com/useresd/ledger/storage/mock"
)

func TestService(t *testing.T) {
	repo := mock.NewCurrency()

	// Populate mock currency data.
	repo.Populate()

	service := currency.NewService(
		log.New(os.Stdout, "TEST", log.Ldate|log.Ltime|log.Lshortfile),
		repo,
	)

	currency, err := service.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "SDG", currency.Abbr)
}

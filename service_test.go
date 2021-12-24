package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	repo := &RepoMock{}
	
	// Populate mock currency data.
	repo.Populate()

	service := &Service{
		log: log.New(os.Stdout, "TEST", log.Ldate|log.Ltime|log.Lshortfile),
		repo: repo,
	}

	currency, err := service.FindByID(1)

	assert.NoError(t, err);
	assert.Equal(t, "SDG", currency.Abbr)
}
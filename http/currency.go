package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/useresd/ledger/app/currency"
	"github.com/useresd/ledger/pkg/rest"
)

type Currency struct {
	Service *currency.Service
}

func NewCurrency(service *currency.Service) *Currency {
	return &Currency{Service: service}
}

func (a *Currency) FindByID(rw http.ResponseWriter, r *http.Request) error {

	// Get the id from route params and convert it to int
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return err
	}

	// Get the currency
	currency, err := a.Service.FindByID(id)
	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, currency)

	// return nil error
	return nil
}

func (a *Currency) GetAll(rw http.ResponseWriter, r *http.Request) error {

	// Get the id from route params and convert it to int
	currencies, err := a.Service.GetAll()
	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, currencies)

	// return nil error
	return nil
}

func (a *Currency) Store(rw http.ResponseWriter, r *http.Request) error {

	body := currency.EditDTO{}

	json.NewDecoder(r.Body).Decode(&body)

	currency, err := a.Service.Store(body)

	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, currency)
	// return nil error
	return nil
}

// Update handles http request to update the currency info
func (a *Currency) Update(rw http.ResponseWriter, r *http.Request) error {
	// Get the id
	id, err := rest.GetID(r)
	if err != nil {
		return err
	}

	// Decode the request body to json
	body := currency.EditDTO{}
	json.NewDecoder(r.Body).Decode(&body)

	// Update the currency
	currency, err := a.Service.Update(id, body)
	if err != nil {
		return err
	}

	// Send the response
	rest.WriteJson(rw, http.StatusOK, currency)
	return nil
}

func (a *Currency) Delete(rw http.ResponseWriter, r *http.Request) error {
	// Get the id
	id, err := rest.GetID(r)
	if err != nil {
		return err
	}

	if err := a.Service.Delete(id); err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, map[string]string{
		"message": "currency deleted",
	})
	return nil

}

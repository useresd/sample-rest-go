package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/useresd/ledger/app/account"
	"github.com/useresd/ledger/pkg/rest"
)

type Account struct {
	Service *account.Service
}

func NewAccount(service *account.Service) *Account {
	return &Account{Service: service}
}

func (a *Account) FindByID(rw http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return err
	}
	account, err := a.Service.FindByID(id)
	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, account)
	return nil
}

func (a *Account) GetAll(rw http.ResponseWriter, r *http.Request) error {
	accounts, err := a.Service.GetAll()
	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, accounts)
	return nil
}

func (a *Account) Store(rw http.ResponseWriter, r *http.Request) error {

	body := account.EditDTO{}

	json.NewDecoder(r.Body).Decode(&body)

	acc, err := a.Service.Store(body)

	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, map[string]interface{}{
		"message": "account created",
		"record":  acc,
	})
	// return nil error
	return nil
}

func (a *Account) Update(rw http.ResponseWriter, r *http.Request) error {

	// get the id to update.
	id, err := rest.GetID(r)
	if err != nil {
		return err
	}

	// get the body
	body := account.EditDTO{}
	json.NewDecoder(r.Body).Decode(&body)

	// pass data to the service.
	updated, err := a.Service.Update(id, body)
	if err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, map[string]interface{}{
		"message": "account updated",
		"record":  updated,
	})
	return nil
}

func (a *Account) Delete(rw http.ResponseWriter, r *http.Request) error {
	// Get the id
	id, err := rest.GetID(r)
	if err != nil {
		return err
	}

	if err := a.Service.Delete(id); err != nil {
		return err
	}

	rest.WriteJson(rw, http.StatusOK, map[string]interface{}{
		"message": "account deleted",
	})

	return nil
}

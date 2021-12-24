package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct{
	Service *Service
}

func NewAPI(service *Service) *API {
	return &API{Service: service}
}

func (a *API) FindByID(rw http.ResponseWriter, r *http.Request) error {
	
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

	WriteJson(rw, http.StatusOK, currency);

	// return nil error
	return nil
}

func (a *API) GetAll(rw http.ResponseWriter, r *http.Request) error {
	
	// Get the id from route params and convert it to int
	currencies, err := a.Service.GetAll()
	if err != nil {
		return err
	}

	WriteJson(rw, http.StatusOK, currencies);

	// return nil error
	return nil
}

func (a *API) Store(rw http.ResponseWriter, r *http.Request) error {
	
	body := struct{
		Name string `json:"name"`
		Abbr string `json:"abbr"`
	}{}

	json.NewDecoder(r.Body).Decode(&body)

	currency, err := a.Service.Store(struct{Name string; Abbr string}(body))

	if err != nil {
		return err
	}

	WriteJson(rw, http.StatusOK, currency)
	// return nil error
	return nil
}


func WriteJson(rw http.ResponseWriter, statusCode int, content interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(content)
}
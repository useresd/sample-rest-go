package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type httpHandler func(http.ResponseWriter, *http.Request) error

func (fn httpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := fn(rw, r); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

var (
	// Account API
	account *AccountAPI
	
	// Currency API
	currency *API

	// logger
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "HTTP", log.Ldate|log.Ltime|log.Lshortfile)
	account = NewAccountAPI(NewAccountService(logger, NewAccountMock().Populate()))
	currency = NewAPI(NewService(logger, NewMock().Populate()))
}	

func main() {
	r := mux.NewRouter()

	// Currency
	r.Handle("/currency/{id}", httpHandler(currency.FindByID)).Methods(http.MethodGet)
	r.Handle("/currency", httpHandler(currency.GetAll)).Methods(http.MethodGet)
	r.Handle("/currency", httpHandler(currency.Store)).Methods(http.MethodPost)

	// Account
	r.Handle("/account/{id}", httpHandler(account.FindByID))
	r.Handle("/account", httpHandler(account.GetAll))

	http.ListenAndServe(":9099", r)
}
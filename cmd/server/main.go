package main

import (
	"log"
	"net/http"
	"os"

	"github.com/useresd/ledger/app/account"
	"github.com/useresd/ledger/app/currency"
	api "github.com/useresd/ledger/http"
	"github.com/useresd/ledger/pkg/rest"
	"github.com/useresd/ledger/storage/mock"

	"github.com/gorilla/mux"
)

type httpHandler func(http.ResponseWriter, *http.Request) error

func (fn httpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := fn(rw, r); err != nil {
		logger.Println(err.Error())
		rest.WriteJson(rw, http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
}

var (
	logger      *log.Logger
	accountAPI  *api.Account
	currencyAPI *api.Currency
)

func init() {
	logger = log.New(os.Stdout, "HTTP", log.Ldate|log.Ltime|log.Lshortfile)
	accountAPI = api.NewAccount(account.NewService(logger, mock.NewAccount().Populate(), mock.NewCurrency().Populate()))
	currencyAPI = api.NewCurrency(currency.NewService(logger, mock.NewCurrency().Populate()))
}

func main() {
	r := mux.NewRouter()

	// Currency
	r.Handle("/currency/{id}", httpHandler(currencyAPI.FindByID)).Methods(http.MethodGet)
	r.Handle("/currency", httpHandler(currencyAPI.GetAll)).Methods(http.MethodGet)
	r.Handle("/currency", httpHandler(currencyAPI.Store)).Methods(http.MethodPost)
	r.Handle("/currency/{id}", httpHandler(currencyAPI.Update)).Methods(http.MethodPut)
	r.Handle("/currency/{id}", httpHandler(currencyAPI.Delete)).Methods(http.MethodDelete)

	// Account
	r.Handle("/account/{id}", httpHandler(accountAPI.FindByID)).Methods(http.MethodGet)
	r.Handle("/account", httpHandler(accountAPI.GetAll)).Methods(http.MethodGet)
	r.Handle("/account", httpHandler(accountAPI.Store)).Methods(http.MethodPost)
	r.Handle("/account/{id}", httpHandler(accountAPI.Update)).Methods(http.MethodPut)
	r.Handle("/account/{id}", httpHandler(accountAPI.Delete)).Methods(http.MethodDelete)

	http.ListenAndServe(":9099", r)
}

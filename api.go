package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gobank/model"
	"gobank/repository"
	"log"
	"net/http"
	"strconv"
)

func WriteJSON(w *http.ResponseWriter, status int, v any) error {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	return json.NewEncoder(*w).Encode(&v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error      string
	repository repository.AccountRepository
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			err := WriteJSON(&writer, http.StatusBadRequest, ApiError{
				Error: err.Error(),
			})
			if err != nil {
				return
			}
		}
	}
}

type ApiServer struct {
	listenAddr string
	repository repository.AccountRepository
}

func NewAPIServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		repository: repository.NewAccountRepository(),
	}
}

func (server *ApiServer) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(server.HandleGetAccounts)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(server.HandleGetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(server.HandleDeleteAccount)).Methods("DELETE")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(server.HandleUpdateAccount)).Methods("PUT")

	log.Fatal(http.ListenAndServe(server.listenAddr, router))
}

func (server *ApiServer) HandleGetAccount(write http.ResponseWriter, request *http.Request) error {
	var err error
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	account, err := server.repository.GetAccountById(id)
	err = WriteJSON(&write, http.StatusOK, account)
	return err
}

func (server *ApiServer) HandleCreateAccount(write http.ResponseWriter, request *http.Request) error {
	var account model.Account
	var err error
	err = json.NewDecoder(request.Body).Decode(&account)
	err = server.repository.CreateAccount(&account)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return WriteJSON(&write, http.StatusCreated, account)
}

func (server *ApiServer) HandleDeleteAccount(writer http.ResponseWriter, request *http.Request) error {

	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	err := server.repository.DeleteAccount(id)
	if err != nil {
		return err
	}

	return WriteJSON(&writer, http.StatusOK, "Account deleted successfully")

}

func (server *ApiServer) HandleTransfer(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *ApiServer) HandleGetAccounts(writer http.ResponseWriter, request *http.Request) error {
	accounts, _ := server.repository.GetAccounts()
	return WriteJSON(&writer, http.StatusOK, accounts)
}

func (server *ApiServer) HandleUpdateAccount(writer http.ResponseWriter, request *http.Request) error {
	var err error
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	var requestUser model.Account
	err = json.NewDecoder(request.Body).Decode(&requestUser)
	if err != nil {
		return err
	}
	err = server.repository.UpdateAccount(id, &requestUser)
	if err != nil {
		return err
	}
	return WriteJSON(&writer, http.StatusOK, "Account Updated Successfully")
}

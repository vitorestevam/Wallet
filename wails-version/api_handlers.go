package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewApi(port string, db database) (api, error) {
	api := api{
		db: db,
	}

	return api, nil
}

type api struct {
	db database
}

func (a *api) PostTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IM IN POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	transaction := Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = a.db.InsertTransaction(&transaction)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(transaction)
	w.Write(response)
}

func (a *api) GetTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IM IN GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	transactions, err := a.db.GetTransactions(nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(transactions)
	w.Write(response)
}

func (a *api) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	transaction, err := a.db.GetTransaction(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(transaction)
	w.Write(response)
}

func (a *api) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	transaction := Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = a.db.UpdateTransaction(id, &transaction)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(transaction)
	w.Write(response)
}

func (a *api) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := a.db.DeleteTransaction(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

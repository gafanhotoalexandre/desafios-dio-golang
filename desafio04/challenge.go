package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Person struct {
	ID        int     `json:"id,omitempty"`
	FirstName string  `json:"firstname,omitempty"`
	LastName  string  `json:"lastname,omitempty"`
	Address   Address `json:"address"`
}

var customers = make(map[int]Person)
var nextID = 1

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /customers", getCustomers)
	mux.HandleFunc("POST /customers", createCustomer)
	mux.HandleFunc("GET /customers/{id}", getCustomer)
	mux.HandleFunc("PUT /customers/{id}", updateCustomer)
	mux.HandleFunc("DELETE /customers/{id}", deleteCustomer)

	fmt.Println("Grandma's Simple Customer API running on :3333")
	log.Fatal(http.ListenAndServe(":3333", mux))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	list := make([]Person, 0, len(customers))

	for _, customer := range customers {
		list = append(list, customer)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Person
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer.ID = nextID
	nextID++
	customers[customer.ID] = customer

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	customer, found := customers[id]
	if !found {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	if _, found := customers[id]; !found {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	var customer Person
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer.ID = id
	customers[id] = customer

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, found := customers[id]; !found {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	delete(customers, id)
	w.WriteHeader(http.StatusNoContent)
}

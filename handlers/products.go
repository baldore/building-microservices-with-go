package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/baldore/building-microservices-with-go/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

type SuccessResponse struct {
	Status string `json:"status"`
}

var sourceCreatedResponse = &SuccessResponse{
	Status: "OK",
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get products")

	pl := data.GetProducts()
	err := pl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle add products")

	np := &data.Product{}
	err := np.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	data.AddProduct(np)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sourceCreatedResponse)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle update products")

	up := &data.Product{}
	if err := up.FromJSON(r.Body); err != nil {
		http.Error(w, "Unable to process data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error: id must be an int", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, up)
	if err != nil {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sourceCreatedResponse)
}

func (p *Products) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle delete products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error: id must be an int", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}
}

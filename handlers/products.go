package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/baldore/building-microservices-with-go/data"
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
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle add products")

	np := &data.Product{}
	err := np.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to process data", http.StatusBadRequest)
	}

	data.AddProduct(np)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sourceCreatedResponse)
}

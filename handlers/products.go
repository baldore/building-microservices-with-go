package handlers

import (
	"log"
	"net/http"

	"github.com/baldore/building-microservices-with-go/data"
)

type Products struct {
	l *log.Logger
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

package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(r io.Writer) error {
	e := json.NewEncoder(r).Encode(p)
	return e
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func GetProducts() Products {
	return productList
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	i, err := findProductIndex(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[i] = p

	return nil
}

func DeleteProduct(id int) error {
	i, err := findProductIndex(id)
	if err != nil {
		return err
	}

	productList = append(productList[:i], productList[i+1:]...)

	return nil
}

func findProductIndex(id int) (int, error) {
	for i := range productList {
		if productList[i].ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

// Hardcoded product list
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

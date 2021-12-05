package data

import (
	"encoding/json"
	"io"
	"time"
)

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

type Products []*Product

func (p *Products) ToJson(rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(p)
}

func GetProduct() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Description: "From milky coffee",
		Name:        "Latte",
		Price:       12.2,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Description: "Short and strong cofee without milk",
		Name:        "Espresso",
		Price:       1.99,
		SKU:         "fjd12",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)

}
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil

}

var ErrorProductNotFOund = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	log.Print(id)
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, 0, ErrorProductNotFOund
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func (p *Product) Validate() error {

	valid := validator.New()
	valid.RegisterValidation("sku", validateSku)
	return valid.Struct(p)
}
func validateSku(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
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

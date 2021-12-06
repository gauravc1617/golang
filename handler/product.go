package handler

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"../data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, rd *http.Request) {
	if rd.Method == http.MethodGet {
		p.getProducts(rw, rd)
		return
	}

	if rd.Method == http.MethodPost {
		p.addProduct(rw, rd)
	}

	if rd.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(rd.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "INvalid Uri", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.updateProduct(id, rw, rd)
		p.l.Print(id)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, rd *http.Request) {
	lp := data.GetProduct()

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Cannot marshall product", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, rd *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(rd.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall to json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	p.l.Printf("Prod %#v", prod)
}
func (p *Products) updateProduct(id int, rw http.ResponseWriter, rd *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(rd.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall to json", http.StatusBadRequest)
	}
	error1 := data.UpdateProduct(id, prod)

	if error1 == data.ErrorProductNotFOund {
		http.Error(rw, "Product nmot found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product nmot found", http.StatusInternalServerError)
	}
}

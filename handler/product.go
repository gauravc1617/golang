package handler

import (
	"log"
	"net/http"

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
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, rd *http.Request) {
	lp := data.GetProduct()

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Cannot marshall product", http.StatusInternalServerError)
	}
}

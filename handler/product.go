package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, rd *http.Request) {
	lp := data.GetProduct()

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Cannot marshall product", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, rd *http.Request) {
	prod := rd.Context().Value(KeyPorduct{}).(*data.Product)
	data.AddProduct(prod)
	p.l.Printf("Prod %#v", prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, rd *http.Request) {
	vars := mux.Vars(rd)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle Put product", id)
	prod := rd.Context().Value((KeyPorduct{})).(*data.Product)
	error1 := data.UpdateProduct(id, prod)

	if error1 == data.ErrorProductNotFOund {
		http.Error(rw, "Product nmot found", http.StatusNotFound)
	}
}

type KeyPorduct struct {
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshall to json", http.StatusBadRequest)
		}

		err1 := prod.Validate()
		if err1 != nil {
			p.l.Print("Error in validating format")
			p.l.Print(err1)

			http.Error(rw, fmt.Sprintf("Error validating %s", err1), http.StatusBadRequest)
			return

		}
		ctx := context.WithValue(r.Context(), KeyPorduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}

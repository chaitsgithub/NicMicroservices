package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Request")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Request")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)

	w.Write([]byte("Record Added Successfully!"))
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle PUT Request")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error converting ID into number", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(prod, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	w.Write([]byte("Record Updated Successfully!"))
}

type KeyProduct struct{}

func (p *Products) GetProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Printf("In GetProductMiddleware!")
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarshall data", http.StatusBadRequest)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), KeyProduct{}, prod))
		next.ServeHTTP(w, r)
	})
}

func (p *Products) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Printf("In LoggingMiddleware!")
		next.ServeHTTP(w, r)
	})
}

package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"net/http"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Request")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)

	w.Write([]byte("Record Added Successfully!"))
}

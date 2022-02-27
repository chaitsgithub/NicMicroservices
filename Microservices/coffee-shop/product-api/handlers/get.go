package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"net/http"
)

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Request")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall products", http.StatusInternalServerError)
	}
}

package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

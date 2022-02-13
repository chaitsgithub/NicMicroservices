package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	if r.Method == http.MethodPut {
		p.updProduct(w, r)
		return
	}

	http.Error(w, "The api only supports GET and POST. Cannot call with "+r.Method, http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Request")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall products", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Request")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall data", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)

	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle PUT Request")
	// Expect the ID in the path
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}
	if len(g[0]) != 2 {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
	}

	idstring := g[0][1]
	id, err := strconv.Atoi(idstring)
	if err != nil {
		http.Error(w, "Error converting ID into number", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall data", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(prod, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	w.Write([]byte("Record Updated Successfully!"))
}

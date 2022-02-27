package handlers

import (
	"NicJackson/Microservices/coffee-shop/product-api/data"
	"context"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p *Products) GetProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Printf("In GetProductMiddleware!")
		prod := &data.Product{}

		if r.Method != http.MethodDelete {
			err := prod.FromJSON(r.Body)
			if err != nil {
				p.l.Printf("Error in GetProductMiddleware : %v\n", err)
				http.Error(w, "Unable to unmarshall data", http.StatusBadRequest)
				return
			}

			err = prod.Validate()
			if err != nil {
				p.l.Printf("Error validating request: %v\n", err)
				http.Error(w, "Error validating request", http.StatusBadRequest)
			}

			r = r.WithContext(context.WithValue(r.Context(), KeyProduct{}, prod))

		}
		next.ServeHTTP(w, r)
	})
}

func (p *Products) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Printf("In LoggingMiddleware!")
		next.ServeHTTP(w, r)
	})
}

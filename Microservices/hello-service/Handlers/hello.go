package Handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Handler!")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops Error!",http.StatusBadRequest)
	}
	
	fmt.Fprintf(w, "Hello %s\n", b)
}


package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required""`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gte=0,lte=130"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w http.ResponseWriter) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", SKUValidator)
	return validate.Struct(p)
}

func SKUValidator(fl validator.FieldLevel) bool {
	r := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := r.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextProductID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

func getNextProductID() int {
	return len(productList) + 1
}

func UpdateProduct(p *Product, id int) error {
	for k, v := range productList {
		if v.ID == id {
			productList[k].Name = p.Name
			productList[k].Description = p.Description
			productList[k].Price = p.Price
			productList[k].SKU = p.SKU
			productList[k].UpdatedOn = time.Now().UTC().String()
			return nil
		}
	}
	return errors.New("Product Not Found for ID : " + strconv.Itoa(id))
}

func DeleteProduct(id int) error {
	for k, v := range productList {
		if v.ID == id {
			productList = append(productList[:k], productList[k+1:]...)
			return nil
		}
	}
	return errors.New("Product Not Found for ID : " + strconv.Itoa(id))
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Expresso",
		Description: "Strong coffee without milk",
		Price:       1.99,
		SKU:         "def456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

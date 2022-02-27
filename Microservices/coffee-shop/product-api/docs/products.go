package docs

import "NicJackson/Microservices/coffee-shop/product-api/data"

// A list of products returns in the response
// swagger:response ProductsResponse
type ProductsResponseWrapper struct {
	// All products in the system
	// in: Body
	Products []data.Product
}

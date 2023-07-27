package docs

import "github.com/joshuabl97/product-api/data"

// A list of products
// swagger:response productsResponse
type productsResponsWrapper struct {
	// all products in the data store
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct listProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct {
	// The requested product information
	// in: body
	Body data.Product
}

// swagger:response noContent
type noContent struct{}

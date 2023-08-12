package ports

import "web-api/internal/core/domain"

type ProductService interface {
	GetAllProducts() ([]*domain.Product, error)
	//GetProductById(Id string) domain.Product
	//SaveProducts(product domain.Product) bool
}
type ProductRepository interface {
	GetAllProducts() ([]*domain.Product, error)
	//GetProductById(Id string) domain.Product
	//SaveProducts(product domain.Product) bool
}

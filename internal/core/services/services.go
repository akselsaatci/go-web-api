package services

import (
	"web-api/internal/core/domain"
	"web-api/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
func (p *ProductService) GetAllProducts() ([]*domain.Product, error) {
	return p.repo.GetAllProducts()
}

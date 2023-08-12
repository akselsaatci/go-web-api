package repository

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"web-api/internal/core/domain"

	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	db *gorm.DB
}

func NewProductPostgresRepository() *ProductPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "123456"
	dbname := "ticaret"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		return nil
	}

	return &ProductPostgresRepository{
		db: db,
	}
}

func (p *ProductPostgresRepository) GetAllProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	req := p.db.Find(&products)
	fmt.Println(products)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return products, nil
}

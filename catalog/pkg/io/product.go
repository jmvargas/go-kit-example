package io

import (
	"errors"
)

type Product struct {
	Id    int32   `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type ProductRepository interface {
	All() ([]Product, error)
	Get(id int32) (Product, error)
}

type basicProductRepository struct {
	data map[int32]Product
}

func (r *basicProductRepository) All() ([]Product, error) {
	data := []Product{}
	for _, p := range r.data {
		data = append(data, p)
	}
	return data, nil
}

func (r *basicProductRepository) Get(id int32) (Product, error) {
	product, ok := r.data[id]
	if !ok {
		return Product{}, errors.New("Product not found")
	}
	return product, nil
}

func NewProductRepository() ProductRepository {
	return &basicProductRepository{
		data: map[int32]Product{
			1: Product{
				Id:    1,
				Title: "Product 1",
				Price: 20.6,
			},
			2: Product{
				Id:    2,
				Title: "Product 2",
				Price: 20.6,
			},
			3: Product{
				Id:    3,
				Title: "Product 3",
				Price: 20.6,
			},
		},
	}
}

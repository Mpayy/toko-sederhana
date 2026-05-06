package repository

import (
	"context"
	"fmt"
	"testing"
	toko_sederhana "toko-sederhana"
	"toko-sederhana/entity"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductRepositoryImpl_TambahProduct(t *testing.T) {
	productRepository := NewProductRepositoryImpl(toko_sederhana.GetConnection())
	ctx := context.Background()

	newProduct := entity.Products{
		Name:  "Sampoerna Mild",
		Price: 34000,
		Stock: 100,
	}

	product, err := productRepository.TambahProduct(ctx, newProduct)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(product)
}

func TestProductRepositoryImpl_GetAllProducts(t *testing.T) {
	productRepository := NewProductRepositoryImpl(toko_sederhana.GetConnection())
	ctx := context.Background()

	products, err := productRepository.GetAllProducts(ctx)
	if err != nil {
		t.Error(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}

//func TestProductRepositoryImpl_GetProductById(t *testing.T) {
//	productRepository := NewProductRepositoryImpl(toko_sederhana.GetConnection())
//	ctx := context.Background()
//	productId := int32(2)
//	product, err := productRepository.GetProductById(ctx, productId)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(product)
//}

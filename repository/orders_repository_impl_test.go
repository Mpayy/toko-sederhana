package repository

import (
	"context"
	"fmt"
	"testing"
	toko_sederhana "toko-sederhana"
	"toko-sederhana/entity"

	_ "github.com/go-sql-driver/mysql"
)

func TestOrderRepositoryImpl_BuatOrder(t *testing.T) {
	db := toko_sederhana.GetConnection()
	productRepository := NewProductRepositoryImpl(db)
	orderRepository := NewOrderRepositoryImpl(db)

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	productIdUpdate := int32(2)
	product, err := productRepository.GetByIDTx(ctx, tx, productIdUpdate)
	if err != nil {
		t.Fatal(err)
	}

	qty := int32(2)
	newStock := product.Stock - qty

	newOrder := entity.Orders{
		ProductId:  product.Id,
		Quantity:   qty,
		TotalPrice: product.Price * qty,
	}

	if product.Stock < newOrder.Quantity {
		t.Fatal("product stock less than qty")
	}

	order, err := orderRepository.BuatOrder(ctx, tx, newOrder)
	if err != nil {
		t.Fatal(err)
	}

	err = productRepository.UpdateStockTx(ctx, tx, productIdUpdate, newStock)
	if err != nil {
		t.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(order)
}

func TestOrderRepositoryImpl_GetAllOrders(t *testing.T) {
	orderRepository := NewOrderRepositoryImpl(toko_sederhana.GetConnection())
	ctx := context.Background()

	orders, err := orderRepository.GetAllOrders(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, order := range orders {
		fmt.Println(order)
	}
}

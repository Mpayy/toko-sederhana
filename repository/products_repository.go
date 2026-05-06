package repository

import (
	"context"
	"database/sql"
	"toko-sederhana/entity"
)

type ProductsRepository interface {
	TambahProduct(ctx context.Context, product entity.Products) (entity.Products, error)
	GetAllProducts(ctx context.Context) ([]entity.Products, error)
	GetByIDTx(ctx context.Context, tx *sql.Tx, productId int32) (entity.Products, error)
	UpdateStockTx(ctx context.Context, tx *sql.Tx, productId int32, newStock int32) error
}

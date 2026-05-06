package repository

import (
	"context"
	"database/sql"
	"toko-sederhana/entity"
)

type OrdersRepository interface {
	BuatOrder(ctx context.Context, tx *sql.Tx, order entity.Orders) (entity.Orders, error)
	GetAllOrders(ctx context.Context) ([]entity.OrderResponse, error)
}

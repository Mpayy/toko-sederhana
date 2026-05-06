package repository

import (
	"context"
	"database/sql"
	"toko-sederhana/entity"
)

type orderRepositoryImpl struct {
	DB *sql.DB
}

func NewOrderRepositoryImpl(db *sql.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{DB: db}
}

func (repository *orderRepositoryImpl) BuatOrder(ctx context.Context, tx *sql.Tx, order entity.Orders) (entity.Orders, error) {
	query := "INSERT INTO orders(product_id, quantity, total_price) VALUES (?, ?, ?)"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return order, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, order.ProductId, order.Quantity, order.TotalPrice)
	if err != nil {
		return order, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return order, err
	}

	order.Id = int32(id)
	return order, nil
}

func (repository *orderRepositoryImpl) GetAllOrders(ctx context.Context) ([]entity.OrderResponse, error) {
	query := "SELECT orders.id, products.name AS product_name, orders.quantity, orders.total_price, orders.created_at FROM orders JOIN products ON products.id = orders.product_id"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []entity.OrderResponse
	for rows.Next() {
		order := entity.OrderResponse{}
		err := rows.Scan(&order.Id, &order.ProductName, &order.Quantity, &order.TotalPrice, &order.CreatedAt)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

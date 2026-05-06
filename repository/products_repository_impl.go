package repository

import (
	"context"
	"database/sql"
	"errors"
	"toko-sederhana/entity"
)

type productRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) *productRepositoryImpl {
	return &productRepositoryImpl{DB: db}
}

func (repository *productRepositoryImpl) TambahProduct(ctx context.Context, product entity.Products) (entity.Products, error) {
	//TODO implement me
	query := "INSERT INTO products (name, price, stock) VALUES (?,?,?)"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	if err != nil {
		return product, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Stock)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}

	product.Id = int32(id)
	return product, nil
}

func (repository *productRepositoryImpl) GetAllProducts(ctx context.Context) ([]entity.Products, error) {
	//TODO implement me
	query := "SELECT id, name, price, stock FROM products"
	//stmt, err := repository.DB.PrepareContext(ctx, query)
	//if err != nil {
	//	return nil, err
	//}
	//defer stmt.Close()

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Products
	for rows.Next() {
		product := entity.Products{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repository *productRepositoryImpl) GetByIDTx(ctx context.Context, tx *sql.Tx, productId int32) (entity.Products, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = ?"
	//row := repository.DB.QueryRowContext(ctx, query, productId)
	//var product entity.Products
	//err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
	//if err != nil {
	//	return product, err
	//}
	//return product, nil

	product := entity.Products{}
	rows := tx.QueryRowContext(ctx, query, productId)
	err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repository *productRepositoryImpl) UpdateStockTx(ctx context.Context, tx *sql.Tx, productId int32, newStock int32) error {
	query := "UPDATE products SET stock = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, newStock, productId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("product not found")
	}
	return nil
}

package toko_sederhana

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/toko-sederhana")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

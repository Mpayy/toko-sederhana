package toko_sederhana

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}
}

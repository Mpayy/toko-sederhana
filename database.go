package toko_sederhana

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
    // Mencoba mencari .env di folder saat ini, jika tidak ada coba di folder atasnya
    // Ini berguna agar saat dijalankan dari test (subfolder) tetap bisa baca .env di root
    err := godotenv.Load(".env")
    if err != nil {
        // Jika gagal, coba cari di satu tingkat di atasnya (../.env)
        err = godotenv.Load("../.env")
        if err != nil {
            fmt.Println("No .env file found in current or parent directory")
        }
    }
}

func GetConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

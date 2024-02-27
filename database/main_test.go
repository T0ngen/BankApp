package database

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	connStr := "user=root password=1234 dbname=bankdbtest2 sslmode=disable port=5433"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Ошибка при подключении к базе данных: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	code := m.Run()

	os.Exit(code)

}



package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=Dhruva@27 dbname=postgres sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("sql.Open ERROR:", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB NOT CONNECTED ✅ ERROR BELOW:")
		fmt.Println(err) // <-- IMPORTANT, prints TRUE reason
		return
	}

	fmt.Println("✅ PostgreSQL connected successfully!")
}

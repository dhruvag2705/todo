package models

import (
	"database/sql"		// SQL database package
	"fmt"				// For formatted I/O

	_ "github.com/lib/pq"		// PostgreSQL driver
)

var DB *sql.DB					// Global DB variable

func ConnectDB() {				// PostgreSQL connection
	var err error				// Stores error if any
	connStr := "host=localhost port=5432 user=postgres password=Dhruva@27 dbname=postgres sslmode=disable"
	//setup level error or problem from go to postgres
	DB, err = sql.Open("postgres", connStr)		// Open connection(Setup DB)
	if err != nil {								// Check error
		fmt.Println("sql.Open ERROR:", err)		// Print error
		return
	}
	//connection level error or problem from postgres
	err = DB.Ping()								// Ping DB to check connection
	if err != nil {								// If error
		fmt.Println("DB NOT CONNECTED ✅ ERROR BELOW:")	// Print message
		fmt.Println(err)						// <-- IMPORTANT, prints TRUE reason
		return
	}

	fmt.Println("✅ PostgreSQL connected successfully!")	// Success message
}

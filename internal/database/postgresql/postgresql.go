// database/postregresql/postregresql.go
package postgresql

import (
	"OnlineBar/pkg/cfg"
	"database/sql"
	"fmt"
	"log"

	// Import the PostgreSQL driver
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	config := cfg.DBConfig()
	pqConnStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Database.DBUser, config.Database.DBPass, config.Database.DBAddr,
		config.Database.DBPort, config.Database.DBName)

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", pqConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check the connection.
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Database connected!")
}

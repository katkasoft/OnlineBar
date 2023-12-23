// database/postregresql/postregresql.go
package postgresql

import (
	"OnlineBar/Backend/pkg/cfg"
	"database/sql"
	"fmt"
	"log"

	// Import the PostgreSQL driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	config := cfg.DBConfig()
	pqConnStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Database.DBUser, config.Database.DBPass, config.Database.DBAddr,
		config.Database.DBPort, config.Database.DBName)

	var err error
	db, err = sql.Open("postgres", pqConnStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Database connected!")
}

func AddUser(name string, email string, password string, os string) error {
	id := uuid.New()
	_, err := db.Query("INSERT INTO \"User\" (id, name, email, password, os) VALUES ($1, $2, $3, $4, $5)", id, name, email, password, os)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UserExist(name string, email string) (error, bool) {

	rows, err := db.Query("SELECT name, email FROM \"User\" WHERE name = $1 AND email = $2", name, email)
	if err != nil {
		log.Println(err)
		return err, false
	}
	defer rows.Close()

	if rows.Next() {
		return nil, true
	}

	return nil, false
}

func GetUserPassword(name string) (string, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM \"User\" WHERE name = $1", name).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
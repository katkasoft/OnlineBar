// database/postregresql/postregresql.go
package postgresql

import (
	"OnlineBar/Backend/internal/models"
	"OnlineBar/Backend/pkg/cfg"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

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

func StartTransaction() (*sql.Tx, error) {
	tx, err := db.Begin()
	return tx, err

}

func EndTransaction(tx *sql.Tx) error {

	if err := tx.Commit(); err != nil {

		return err

	}

	return nil
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

func GetUserID(name string) (string, error) {
	var id string
	err := db.QueryRow("SELECT id FROM \"User\" WHERE name = $1", name).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil

}

func PostBuyList(tx *sql.Tx, userID string, product string, price float64, quantity float64, date time.Time) error {
	var balance float64

	_, err := tx.Exec(`
		INSERT INTO BuyList (userID, name, price, quantity, date)
		VALUES ($1, $2, $3, $4, $5)`, userID, product, price, quantity, date)

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE \"User\" SET balance = balance - $1 WHERE id = $2", price, userID)

	if err != nil {
		return err
	}

	err = tx.QueryRow("SELECT balance FROM \"User\" WHERE id = $1", userID).Scan(&balance)
	if balance < 0 || err != nil {
		log.Println(balance)
		log.Println("Don't have enough money")
		return errors.New("don't have enough money")
	}

	return nil
}

func GetBuyList(userID string) (models.ProductList, error) {
	var productList models.ProductList

	rows, err := db.Query("SELECT name, price, quantity, date FROM buylist WHERE userid = $1", userID)
	if err != nil {
		return productList, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Name, &product.Cost, &product.Quantity, &product.Data)
		if err != nil {
			return productList, err
		}
		productList.Products = append(productList.Products, product)
	}

	if err := rows.Err(); err != nil {
		return productList, err
	}

	return productList, nil
}

func GetBalance(userID string) (float64, error) {
	var balance float64

	err := db.QueryRow("SELECT balance FROM \"User\" WHERE id = $1", userID).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func UpdateBalance(balance float64, userID string) error {

	_, err := db.Query("UPDATE \"User\" SET balance = $1 WHERE id = $2", balance, userID)
	if err != nil {
		return err
	}

	return nil
}

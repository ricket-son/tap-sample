package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	dbUsername, err := getSecret("username")
	if err != nil {
		fmt.Printf("Error reading secret: %s\n", err)
	}
	dbPassword, err := getSecret("password")
	if err != nil {
		fmt.Printf("Error reading secret: %s\n", err)
	}
	dbHost, err := getSecret("host")
	if err != nil {
		fmt.Printf("Error reading secret: %s\n", err)
	}
	dbPort, err := getSecret("port")
	if err != nil {
		fmt.Printf("Error reading secret: %s\n", err)
	}
	dbName, err := getSecret("database")
	if err != nil {
		fmt.Printf("Error reading secret: %s\n", err)
	}

	// Create a database connection string
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertTablesIfNotExists(db)

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, db)
		case http.MethodPost:
			addItem(w, r, db)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	port := 8080
	fmt.Printf("Server listening on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getSecret(name string) (string, error) {
	filePath := os.Getenv("SERVICE_BINDING_ROOT") + "/db/" + name
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %w", filePath, err)
	}
	return string(dat), nil
}

func insertTablesIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			price INT
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func getItems(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, price FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO items (name, price) VALUES (?, ?)", newItem.Name, newItem.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

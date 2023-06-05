package customers_who_bought_all_products

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCustomersWhoBoughtAllProducts(t *testing.T) {
	// Read SQL query from the file.
	query, err := ioutil.ReadFile("solution.sql")
	if err != nil {
		t.Fatalf("Failed to read SQL query from solution.sql: %s", err)
	}

	// Connect to the database.
	db, err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	// Drop the Customer and Product table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Customer, Product`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Customer and Product table.
	_, err = db.Exec(`CREATE TABLE Customer (
		customer_id int,
		product_key int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Customer table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Product (
		product_key int PRIMARY KEY
	)`)

	if err != nil {
		t.Fatalf("Failed to create Product table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Customer (customer_id, product_key) VALUES 
		(1, 5), 
		(2, 6),
		(3, 5),
		(3, 6),
		(1, 6)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Customer table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Product (product_key) VALUES 
		(5),
		(6)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Product table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]struct{}{
		1: {},
		3: {},
	}

	for rows.Next() {
		var customer_id int
		err := rows.Scan(&customer_id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		_, ok := expected[customer_id]
		if !ok {
			t.Errorf("Unexpected result for customer_id: %d.", customer_id)
		}

		delete(expected, customer_id)
	}

	if len(expected) != 0 {
		t.Errorf("Not all customers who bought all products were returned")
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

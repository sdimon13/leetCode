package rearrange_products_table

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRearrangeProductsTable(t *testing.T) {
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

	// Drop the Products table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Products`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Products table.
	_, err = db.Exec(`CREATE TABLE Products (
  		product_id int PRIMARY KEY,
  		store1 int,
		store2 int,
		store3 int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Products (product_id, store1, store2, store3) VALUES 
		(0, 95, 100, 105), 
		(1, 70, NULL, 80)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		product_id int
		store      string
		price      int
	}{
		{0, "store1", 95},
		{0, "store2", 100},
		{0, "store3", 105},
		{1, "store1", 70},
		{1, "store3", 80},
	}

	i := 0
	for rows.Next() {
		var product_id int
		var store string
		var price int
		err := rows.Scan(&product_id, &store, &price)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if product_id != expected[i].product_id || store != expected[i].store || price != expected[i].price {
			t.Errorf("Row %d - Expected (product_id: %d, store: %s, price: %d), got (product_id: %d, store: %s, price: %d)",
				i+1, expected[i].product_id, expected[i].store, expected[i].price, product_id, store, price)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

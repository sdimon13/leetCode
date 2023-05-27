package recyclable_and_low_fat_products

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRecyclableAndLowFatProducts(t *testing.T) {
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

	// Create the Products table.
	_, err = db.Exec(`CREATE TABLE Products (
  		product_id int,
  		low_fats enum('Y', 'N'),
  		recyclable enum('Y', 'N')
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Products (product_id, low_fats, recyclable) VALUES 
		(0, 'Y', 'N'), 
		(1, 'Y', 'Y'), 
		(2, 'N', 'Y'), 
		(3, 'Y', 'Y'), 
		(4, 'N', 'N')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	var productIds []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		productIds = append(productIds, id)
	}

	expectedProductIds := []int{1, 3}

	if len(productIds) != len(expectedProductIds) {
		t.Errorf("Expected %d productIds, got %d", len(expectedProductIds), len(productIds))
	}

	for i, id := range productIds {
		if id != expectedProductIds[i] {
			t.Errorf("Expected product id at position %d to be %d, got %d", i, expectedProductIds[i], id)
		}
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Products`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

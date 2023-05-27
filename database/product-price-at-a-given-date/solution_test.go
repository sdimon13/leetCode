package product_price_at_a_given_date

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductPriceAtGivenDate(t *testing.T) {
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
		product_id int,
		new_price int,
		change_date date,
		PRIMARY KEY (product_id, change_date)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Products table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Products (product_id, new_price, change_date) VALUES 
		(1, 20, '2019-08-14'),
		(2, 50, '2019-08-14'),
		(1, 30, '2019-08-15'),
		(1, 35, '2019-08-16'),
		(2, 65, '2019-08-17'),
		(3, 20, '2019-08-18')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Products table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]int{
		1: 35,
		2: 50,
		3: 10,
	}

	for rows.Next() {
		var product_id int
		var price int
		err := rows.Scan(&product_id, &price)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedPrice, ok := expected[product_id]
		if !ok || price != expectedPrice {
			t.Errorf("Unexpected result for product_id: %d. got price: %d, want: %d", product_id, price, expectedPrice)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package sales_analysis_iii

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSalesAnalysisIII(t *testing.T) {
	// Read SQL query from the file.
	query, err := ioutil.ReadFile("solution.sql")
	if err != nil {
		t.Fatalf("Failed to read SQL query from solution.sql: %s", err)
	}

	// Connect to the database.
	db, err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %s", err)
	}
	defer db.Close()

	// Drop the Product and Sales table if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Product, Sales`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Product and Sales table.
	_, err = db.Exec(`CREATE TABLE Product (
		product_id int PRIMARY KEY,
		product_name varchar(255),
		unit_price int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Product table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Sales (
		seller_id int,
		product_id int,
		buyer_id int,
		sale_date date,
		quantity int,
		price int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Sales table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Product (product_id, product_name, unit_price) VALUES 
		(1, 'S8', 1000), 
		(2, 'G4', 800),
		(3, 'iPhone', 1400)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Product table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Sales (seller_id, product_id, buyer_id, sale_date, quantity, price) VALUES 
		(1, 1, 1, '2019-01-21', 2, 2000), 
		(1, 2, 2, '2019-02-17', 1, 800),
		(2, 2, 3, '2019-06-02', 1, 800),
		(3, 3, 4, '2019-05-13', 2, 2800)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Sales table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]string{
		1: "S8",
	}

	for rows.Next() {
		var product_id int
		var product_name string
		err := rows.Scan(&product_id, &product_name)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedProductName, ok := expected[product_id]
		if !ok || product_name != expectedProductName {
			t.Errorf("Unexpected result for product_id: %d. got product_name: %s, want: %s", product_id, product_name, expectedProductName)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close the database connection: %s", err)
	}
}

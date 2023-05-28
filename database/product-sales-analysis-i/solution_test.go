package product_sales_analysis_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductSalesAnalysis(t *testing.T) {
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

	// Drop the Sales and Product tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Sales, Product`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Sales and Product tables.
	_, err = db.Exec(`CREATE TABLE Sales (
		sale_id int PRIMARY KEY,
		product_id int,
		year int,
		quantity int,
		price int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Sales table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Product (
		product_id int PRIMARY KEY,
		product_name varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Product table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Sales (sale_id, product_id, year, quantity, price) VALUES 
		(1, 100, 2008, 10, 5000),
		(2, 100, 2009, 12, 5000),
		(7, 200, 2011, 15, 9000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Sales table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Product (product_id, product_name) VALUES 
		(100, 'Nokia'),
		(200, 'Apple'),
		(300, 'Samsung')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Product table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	type Sale struct {
		productName string
		year        int
		price       int
	}

	expected := []Sale{
		{"Nokia", 2008, 5000},
		{"Nokia", 2009, 5000},
		{"Apple", 2011, 9000},
	}

	i := 0
	for rows.Next() {
		var year, price int
		var product_name string
		err := rows.Scan(&product_name, &year, &price)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		if product_name != expected[i].productName || year != expected[i].year || price != expected[i].price {
			t.Errorf("Unexpected result for index: %d. got product_name: %s, year: %d, price: %d, want product_name: %s, year: %d, price: %d", i, product_name, year, price, expected[i].productName, expected[i].year, expected[i].price)
		}

		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

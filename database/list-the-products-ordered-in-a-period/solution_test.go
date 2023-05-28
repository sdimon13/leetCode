package list_the_products_ordered_in_a_period

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductsOrderedInAPeriod(t *testing.T) {
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

	// Drop the Products and Orders table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Products, Orders`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Products and Orders tables.
	_, err = db.Exec(`CREATE TABLE Products (
		product_id int PRIMARY KEY,
		product_name varchar(255),
		product_category varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Products table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Orders (
		product_id int,
		order_date date,
		unit int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Orders table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Products (product_id, product_name, product_category) VALUES 
		(1, 'Leetcode Solutions', 'Book'), 
		(2, 'Jewels of Stringology', 'Book'),
		(3, 'HP', 'Laptop'),
		(4, 'Lenovo', 'Laptop'),
		(5, 'Leetcode Kit', 'T-shirt')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Products table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Orders (product_id, order_date, unit) VALUES 
		(1, '2020-02-05', 60),
		(1, '2020-02-10', 70),
		(2, '2020-01-18', 30),
		(2, '2020-02-11', 80),
		(3, '2020-02-17', 2),
		(3, '2020-02-24', 3),
		(4, '2020-03-01', 20),
		(4, '2020-03-04', 30),
		(4, '2020-03-04', 60),
		(5, '2020-02-25', 50),
		(5, '2020-02-27', 50),
		(5, '2020-03-01', 50)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Orders table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]int{
		"Leetcode Solutions": 130,
		"Leetcode Kit":       100,
	}

	for rows.Next() {
		var product_name string
		var unit int
		err := rows.Scan(&product_name, &unit)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedUnit, ok := expected[product_name]
		if !ok || unit != expectedUnit {
			t.Errorf("Unexpected result for product_name: %s. got unit: %d, want: %d", product_name, unit, expectedUnit)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

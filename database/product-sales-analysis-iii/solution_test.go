package product_sales_analysis_iii

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFirstYearProductSale(t *testing.T) {
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

	// Drop the Sales and Product table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Sales, Product`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Sales and Product table.
	_, err = db.Exec(`CREATE TABLE Sales (
		sale_id int,
		product_id int,
		year int,
		quantity int,
		price int,
		PRIMARY KEY(sale_id, year)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Sales table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Product (
		product_id int PRIMARY KEY,
		product_name varchar(100)
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

	expected := map[int]map[int]struct {
		quantity int
		price    int
	}{
		100: {
			2008: {10, 5000},
		},
		200: {
			2011: {15, 9000},
		},
	}

	for rows.Next() {
		var product_id int
		var first_year int
		var quantity int
		var price int
		err := rows.Scan(&product_id, &first_year, &quantity, &price)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		if _, ok := expected[product_id][first_year]; !ok {
			t.Errorf("Unexpected product_id: %d and first_year: %d", product_id, first_year)
		} else if quantity != expected[product_id][first_year].quantity || price != expected[product_id][first_year].price {
			t.Errorf("Unexpected result for product_id: %d and first_year: %d. got quantity: %d, want: %d and got price: %d, want: %d", product_id, first_year, quantity, expected[product_id][first_year].quantity, price, expected[product_id][first_year].price)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

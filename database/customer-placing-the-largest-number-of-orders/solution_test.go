package customer_placing_the_largest_number_of_orders

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCustomerWithLargestNumberOfOrders(t *testing.T) {
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

	// Drop the Orders table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Orders`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Orders table.
	_, err = db.Exec(`CREATE TABLE Orders (
		order_number int PRIMARY KEY,
		customer_number int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Orders table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Orders (order_number, customer_number) VALUES 
		(1, 1), 
		(2, 2),
		(3, 3),
		(4, 3)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Orders table: %s", err)
	}

	// Execute the SQL query from the file.
	row := db.QueryRow(string(query))

	var customerNumber int
	err = row.Scan(&customerNumber)
	if err != nil {
		t.Fatalf("Failed to scan row: %s", err)
	}

	expectedCustomerNumber := 3

	if customerNumber != expectedCustomerNumber {
		t.Errorf("Unexpected result for customer_number. got: %d, want: %d", customerNumber, expectedCustomerNumber)
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

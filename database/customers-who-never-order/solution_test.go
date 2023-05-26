package customers_who_never_order

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCustomersWhoNeverOrder(t *testing.T) {
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

	// Create the Customers and Orders tables.
	_, err = db.Exec(`CREATE TABLE Customers (id int, name varchar(255))`)
	if err != nil {
		t.Fatalf("Failed to create Customers table: %s", err)
	}
	_, err = db.Exec(`CREATE TABLE Orders (id int, customerId int)`)
	if err != nil {
		t.Fatalf("Failed to create Orders table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Customers (id, name) VALUES (1, 'Joe'), (2, 'Henry'), (3, 'Sam'), (4, 'Max')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Customers table: %s", err)
	}
	_, err = db.Exec(`INSERT INTO Orders (id, customerId) VALUES (1, 3), (2, 1)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Orders table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("SQL query returned error: %s", err)
	}

	var customers []string
	for rows.Next() {
		var customer string
		if err := rows.Scan(&customer); err != nil {
			t.Fatalf("Failed to scan row: %s", err)
		}
		customers = append(customers, customer)
	}

	expectedCustomers := []string{"Henry", "Max"}

	if len(customers) != len(expectedCustomers) {
		t.Errorf("Expected %d customers, got %d", len(expectedCustomers), len(customers))
	}

	sort.Strings(customers)
	sort.Strings(expectedCustomers)

	for i, customer := range customers {
		if customer != expectedCustomers[i] {
			t.Errorf("Expected customer at position %d to be %s, got %s", i, expectedCustomers[i], customer)
		}
	}

	// Drop the tables at the end of the test.
	_, err = db.Exec(`DROP TABLE Customers`)
	if err != nil {
		t.Errorf("Failed to drop Customers table: %s", err)
	}

	_, err = db.Exec(`DROP TABLE Orders`)
	if err != nil {
		t.Errorf("Failed to drop Orders table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

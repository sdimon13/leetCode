package customer_who_visited_but_did_not_make_any_transactions

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCustomerWhoVisitedButDidNotMakeAnyTransactions(t *testing.T) {
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

	// Drop the Visits and Transactions table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Visits, Transactions`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Visits and Transactions table.
	_, err = db.Exec(`CREATE TABLE Visits (
		visit_id int PRIMARY KEY,
		customer_id int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Visits table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Transactions (
		transaction_id int PRIMARY KEY,
		visit_id int,
		amount int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Transactions table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Visits (visit_id, customer_id) VALUES 
		(1, 23), 
		(2, 9),
		(4, 30),
		(5, 54),
		(6, 96),
		(7, 54),
		(8, 54)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Visits table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Transactions (transaction_id, visit_id, amount) VALUES 
		(2, 5, 310), 
		(3, 5, 300),
		(9, 5, 200),
		(12, 1, 910),
		(13, 2, 970)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Transactions table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]int{
		54: 2,
		30: 1,
		96: 1,
	}

	for rows.Next() {
		var customer_id int
		var count_no_trans int
		err := rows.Scan(&customer_id, &count_no_trans)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedCount, ok := expected[customer_id]
		if !ok || count_no_trans != expectedCount {
			t.Errorf("Unexpected result for customer_id: %d. got count_no_trans: %d, want: %d", customer_id, count_no_trans, expectedCount)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

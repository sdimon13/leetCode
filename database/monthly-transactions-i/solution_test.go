package monthly_transactions_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMonthlyTransactions(t *testing.T) {
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

	// Drop the Transactions table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Transactions`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Transactions table.
	_, err = db.Exec(`CREATE TABLE Transactions (
		id int PRIMARY KEY,
		country varchar(50),
		state enum('approved', 'declined'),
		amount int,
		trans_date date
	)`)

	if err != nil {
		t.Fatalf("Failed to create Transactions table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Transactions (id, country, state, amount, trans_date) VALUES 
		(121, 'US', 'approved', 1000, '2018-12-18'), 
		(122, 'US', 'declined', 2000, '2018-12-19'),
		(123, 'US', 'approved', 2000, '2019-01-01'),
		(124, 'DE', 'approved', 2000, '2019-01-07')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Transactions table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]map[string][]int{
		"2018-12": {
			"US": {2, 1, 3000, 1000},
		},
		"2019-01": {
			"US": {1, 1, 2000, 2000},
			"DE": {1, 1, 2000, 2000},
		},
	}

	for rows.Next() {
		var month string
		var country string
		var trans_count, approved_count, trans_total_amount, approved_total_amount int
		err := rows.Scan(&month, &country, &trans_count, &approved_count, &trans_total_amount, &approved_total_amount)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedData, ok := expected[month][country]
		if !ok || trans_count != expectedData[0] || approved_count != expectedData[1] || trans_total_amount != expectedData[2] || approved_total_amount != expectedData[3] {
			t.Errorf("Unexpected result for month: %s and country: %s. got: %d, %d, %d, %d, want: %d, %d, %d, %d",
				month, country, trans_count, approved_count, trans_total_amount, approved_total_amount,
				expectedData[0], expectedData[1], expectedData[2], expectedData[3])
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

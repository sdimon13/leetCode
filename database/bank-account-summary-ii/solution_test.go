package bank_account_summary_ii

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBankAccountSummaryII(t *testing.T) {
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

	// Drop the Users and Transactions table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Users, Transactions`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Users and Transactions table.
	_, err = db.Exec(`CREATE TABLE Users (
		account int PRIMARY KEY,
		name varchar(50)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Transactions (
		trans_id int PRIMARY KEY,
		account int,
		amount int,
		transacted_on date
	)`)

	if err != nil {
		t.Fatalf("Failed to create Transactions table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Users (account, name) VALUES 
		(900001, 'Alice'), 
		(900002, 'Bob'),
		(900003, 'Charlie')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Transactions (trans_id, account, amount, transacted_on) VALUES 
		(1, 900001, 7000, '2020-08-01'), 
		(2, 900001, 7000, '2020-09-01'),
		(3, 900001, -3000, '2020-09-02'),
		(4, 900002, 1000, '2020-09-12'),
		(5, 900003, 6000, '2020-08-07'),
		(6, 900003, 6000, '2020-09-07'),
		(7, 900003, -4000, '2020-09-11')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Transactions table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]int{
		"Alice": 11000,
	}

	for rows.Next() {
		var name string
		var balance int
		err := rows.Scan(&name, &balance)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedBalance, ok := expected[name]
		if !ok || balance != expectedBalance {
			t.Errorf("Unexpected result for name: %s. got balance: %d, want: %d", name, balance, expectedBalance)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

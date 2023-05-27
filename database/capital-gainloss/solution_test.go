package capital_gainloss

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCapitalGainLoss(t *testing.T) {
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

	// Drop the Stocks table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Stocks`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Stocks table.
	_, err = db.Exec(`CREATE TABLE Stocks (
		stock_name varchar(50),
		operation enum('Buy', 'Sell'),
		operation_day int,
		price int,
		PRIMARY KEY (stock_name, operation_day)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Stocks table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Stocks (stock_name, operation, operation_day, price) VALUES 
		('Leetcode', 'Buy', 1, 1000), 
		('Corona Masks', 'Buy', 2, 10),
		('Leetcode', 'Sell', 5, 9000),
		('Handbags', 'Buy', 17, 30000),
		('Corona Masks', 'Sell', 3, 1010),
		('Corona Masks', 'Buy', 4, 1000),
		('Corona Masks', 'Sell', 5, 500),
		('Corona Masks', 'Buy', 6, 1000),
		('Handbags', 'Sell', 29, 7000),
		('Corona Masks', 'Sell', 10, 10000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Stocks table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]int{
		"Corona Masks": 9500,
		"Leetcode":     8000,
		"Handbags":     -23000,
	}

	for rows.Next() {
		var stock_name string
		var capital_gain_loss int
		err := rows.Scan(&stock_name, &capital_gain_loss)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedGainLoss, ok := expected[stock_name]
		if !ok || capital_gain_loss != expectedGainLoss {
			t.Errorf("Unexpected result for stock_name: %s. got capital_gain_loss: %d, want: %d", stock_name, capital_gain_loss, expectedGainLoss)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

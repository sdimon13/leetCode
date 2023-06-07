package count_salary_categories

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCountSalaryCategories(t *testing.T) {
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

	// Drop the Accounts table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Accounts`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Accounts table.
	_, err = db.Exec(`CREATE TABLE Accounts (
		account_id int PRIMARY KEY,
		income int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Accounts table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Accounts (account_id, income) VALUES 
		(3, 108939), 
		(2, 12747),
		(8, 87709),
		(6, 91796)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Accounts table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]int{
		"Low Salary":     1,
		"Average Salary": 0,
		"High Salary":    3,
	}

	for rows.Next() {
		var category string
		var count int
		err := rows.Scan(&category, &count)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedCount, ok := expected[category]
		if !ok || count != expectedCount {
			t.Errorf("Unexpected result for category: %s. got count: %d, want: %d", category, count, expectedCount)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

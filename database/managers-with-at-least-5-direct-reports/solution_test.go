package managers_with_at_least_5_direct_reports

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestManagersWithAtLeastFiveDirectReports(t *testing.T) {
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

	// Drop the Employee table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employee`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employee table.
	_, err = db.Exec(`CREATE TABLE Employee (
		id int PRIMARY KEY,
		name varchar(255),
		department varchar(255),
		managerId int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employee table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employee (id, name, department, managerId) VALUES 
		(101, 'John', 'A', NULL), 
		(102, 'Dan', 'A', 101),
		(103, 'James', 'A', 101),
		(104, 'Amy', 'A', 101),
		(105, 'Anne', 'A', 101),
		(106, 'Ron', 'B', 101)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employee table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := "John"
	var result string

	if rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			t.Fatalf("Failed to scan row: %s", err)
		}
	}

	if result != expected {
		t.Fatalf("Unexpected result. got: %s, want: %s", result, expected)
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package second_highest_salary

import (
	"database/sql"
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	Department string
	Employee   string
	Salary     int
}

func TestSecondHighestSalary(t *testing.T) {
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

	// Create the Employee table.
	_, err = db.Exec(`CREATE TABLE Employee (
  		Id int,
  		Salary int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employee (Id, Salary) VALUES (1, 100), (2, 200), (3, 300)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	row := db.QueryRow(string(query))
	var SecondHighestSalary sql.NullInt64
	err = row.Scan(&SecondHighestSalary)
	if err != nil {
		t.Fatalf("Failed to scan row: %s", err)
	}

	// Define the expected value.
	expectedSalary := sql.NullInt64{Int64: 200, Valid: true}

	if !SecondHighestSalary.Valid && expectedSalary.Valid {
		t.Errorf("Expected second highest salary to be %d, got NULL", expectedSalary.Int64)
	} else if SecondHighestSalary.Valid && !expectedSalary.Valid {
		t.Errorf("Expected second highest salary to be NULL, got %d", SecondHighestSalary.Int64)
	} else if SecondHighestSalary.Int64 != expectedSalary.Int64 {
		t.Errorf("Expected second highest salary to be %d, got %d", expectedSalary.Int64, SecondHighestSalary.Int64)
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Employee`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

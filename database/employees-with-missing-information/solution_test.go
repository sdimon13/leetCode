package employees_with_missing_information

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmployeesWithMissingInformation(t *testing.T) {
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

	// Drop the Employees and Salaries tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employees`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS Salaries`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employees and Salaries tables.
	_, err = db.Exec(`CREATE TABLE Employees (
  		employee_id int PRIMARY KEY,
  		name varchar(100)
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Salaries (
  		employee_id int PRIMARY KEY,
		salary int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Employees (employee_id, name) VALUES 
		(2, 'Crew'), 
		(4, 'Haven'),
		(5, 'Kristian')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Salaries (employee_id, salary) VALUES 
		(5, 76071),
		(1, 22517),
		(4, 63539)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []int{1, 2}

	i := 0
	for rows.Next() {
		var employee_id int
		err := rows.Scan(&employee_id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if employee_id != expected[i] {
			t.Errorf("Expected row %d to be %d, got %d", i+1, expected[i], employee_id)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

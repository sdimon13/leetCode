package primary_department_for_each_employee

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPrimaryDepartment(t *testing.T) {
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
		employee_id int,
		department_id int,
		primary_flag varchar(1),
		PRIMARY KEY (employee_id, department_id)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employee table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employee (employee_id, department_id, primary_flag) VALUES 
		(1, 1, 'N'), 
		(2, 1, 'Y'),
		(2, 2, 'N'),
		(3, 3, 'N'),
		(4, 2, 'N'),
		(4, 3, 'Y'),
		(4, 4, 'N')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employee table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]int{
		1: 1,
		2: 1,
		3: 3,
		4: 3,
	}

	for rows.Next() {
		var employee_id int
		var department_id int
		err := rows.Scan(&employee_id, &department_id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedDepartment, ok := expected[employee_id]
		if !ok || department_id != expectedDepartment {
			t.Errorf("Unexpected result for employee id: %d. got department id: %d, want: %d", employee_id, department_id, expectedDepartment)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

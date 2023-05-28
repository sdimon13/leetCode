package employees_whose_manager_left_the_company

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmployeesWhoseManagerLeftTheCompany(t *testing.T) {
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

	// Drop the Employees table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employees`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employees table.
	_, err = db.Exec(`CREATE TABLE Employees (
		employee_id int PRIMARY KEY,
		name varchar(255),
		manager_id int,
		salary int
	)`)
	if err != nil {
		t.Fatalf("Failed to create Employees table: %s", err)
	}

	// Populate the Employees table with initial data.
	_, err = db.Exec(`INSERT INTO Employees (employee_id, name, manager_id, salary) VALUES 
		(3, 'Mila', 9, 60301), 
		(12, 'Antonella', NULL, 31000), 
		(13, 'Emery', NULL, 67084), 
		(1, 'Kalel', 11, 21241), 
		(9, 'Mikaela', NULL, 50937), 
		(11, 'Joziah', 6, 28485)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Employees table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		employee_id int
	}{
		{11},
	}

	i := 0
	for rows.Next() {
		var employee_id int
		err := rows.Scan(&employee_id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if employee_id != expected[i].employee_id {
			t.Errorf("Row %d - Expected employee_id: %d, got employee_id: %d",
				i+1, expected[i].employee_id, employee_id)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

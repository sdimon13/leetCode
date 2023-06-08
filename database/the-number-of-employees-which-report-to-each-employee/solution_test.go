package the_number_of_employees_which_report_to_each_employee

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestReportToEachEmployee(t *testing.T) {
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

	// Drop the Employees table if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employees`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employees table.
	_, err = db.Exec(`CREATE TABLE Employees (
		employee_id int PRIMARY KEY,
		name varchar(255),
		reports_to int,
		age int
	)`)
	if err != nil {
		t.Fatalf("Failed to create Employees table: %s", err)
	}

	// Populate the Employees table with initial data.
	_, err = db.Exec(`INSERT INTO Employees (employee_id, name, reports_to, age) VALUES 
		(9, 'Hercy', NULL, 43), 
		(6, 'Alice', 9, 41), 
		(4, 'Bob', 9, 36),
		(2, 'Winston', NULL, 37)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Employees table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		employee_id   int
		name          string
		reports_count int
		average_age   int
	}{
		{9, "Hercy", 2, 39},
	}

	i := 0
	for rows.Next() {
		var employee_id, reports_count, average_age int
		var name string
		err := rows.Scan(&employee_id, &name, &reports_count, &average_age)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if employee_id != expected[i].employee_id || name != expected[i].name || reports_count != expected[i].reports_count || average_age != expected[i].average_age {
			t.Errorf("Row %d - Expected (employee_id: %d, name: %s, reports_count: %d, average_age: %d), got (employee_id: %d, name: %s, reports_count: %d, average_age: %d)",
				i+1, expected[i].employee_id, expected[i].name, expected[i].reports_count, expected[i].average_age, employee_id, name, reports_count, average_age)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

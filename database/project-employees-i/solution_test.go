package project_employees_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAverageExperienceYearsOfEmployeesPerProject(t *testing.T) {
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

	// Drop the Project and Employee table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Project, Employee`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Project and Employee table.
	_, err = db.Exec(`CREATE TABLE Project (
		project_id int,
		employee_id int,
		PRIMARY KEY (project_id, employee_id)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Project table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Employee (
		employee_id int PRIMARY KEY,
		name varchar(255),
		experience_years int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employee table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Project (project_id, employee_id) VALUES 
		(1, 1), 
		(1, 2),
		(1, 3),
		(2, 1),
		(2, 4)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Project table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Employee (employee_id, name, experience_years) VALUES 
		(1, 'Khaled', 3), 
		(2, 'Ali', 2),
		(3, 'John', 1),
		(4, 'Doe', 2)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employee table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]float64{
		1: 2.00,
		2: 2.50,
	}

	for rows.Next() {
		var project_id int
		var average_years float64
		err := rows.Scan(&project_id, &average_years)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedAvg, ok := expected[project_id]
		if !ok || math.Abs(average_years-expectedAvg) > 0.01 {
			t.Errorf("Unexpected result for project_id: %d. got average_years: %.2f, want: %.2f", project_id, average_years, expectedAvg)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

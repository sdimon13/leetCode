package department_top_three_salaries

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	Department string
	Employee   string
	Salary     int
}

func TestDepartmentTopThreeSalaries(t *testing.T) {
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

	// Create the Employee and Department tables.
	_, err = db.Exec(`CREATE TABLE Employee (
  		id int,
  		name varchar(255),
  		salary int,
  		departmentId int
	)`)
	if err != nil {
		t.Fatalf("Failed to create Employee table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Department (
  		id int,
  		name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Department table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Employee (id, name, salary, departmentId) VALUES 
		(1, 'Joe', 85000, 1), 
		(2, 'Henry', 80000, 2), 
		(3, 'Sam', 60000, 2),
		(4, 'Max', 90000, 1),
		(5, 'Janet', 69000, 1),
		(6, 'Randy', 85000, 1),
		(7, 'Will', 70000, 1)`)

	_, err = db.Exec(`INSERT INTO Department (id, name) VALUES 
		(1, 'IT'), 
		(2, 'Sales')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}
	defer rows.Close()

	var results []Result
	for rows.Next() {
		var result Result
		if err := rows.Scan(&result.Department, &result.Employee, &result.Salary); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		results = append(results, result)
	}

	expectedResults := []Result{
		{"IT", "Max", 90000},
		{"IT", "Joe", 85000},
		{"IT", "Randy", 85000},
		{"IT", "Will", 70000},
		{"Sales", "Henry", 80000},
		{"Sales", "Sam", 60000},
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Department == results[j].Department {
			return results[i].Salary > results[j].Salary
		}
		return results[i].Department < results[j].Department
	})

	for i, result := range results {
		if result.Department != expectedResults[i].Department ||
			result.Employee != expectedResults[i].Employee ||
			result.Salary != expectedResults[i].Salary {
			t.Errorf("Expected result at position %d to be %+v, got %+v", i, expectedResults[i], result)
		}
	}

	// Drop the tables at the end of the test.
	_, err = db.Exec(`DROP TABLE Employee`)
	if err != nil {
		t.Errorf("Failed to drop Employee table: %s", err)
	}
	_, err = db.Exec(`DROP TABLE Department`)
	if err != nil {
		t.Errorf("Failed to drop Department table: %s", err)
	}
}

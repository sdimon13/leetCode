package department_highest_salary

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDepartmentHighestSalary(t *testing.T) {
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
	_, err = db.Exec(`INSERT INTO Employee (id, name, salary, departmentId) VALUES (1, 'Joe', 70000, 1), (2, 'Jim', 90000, 1), (3, 'Henry', 80000, 2), (4, 'Sam', 60000, 2), (5, 'Max', 90000, 1)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Employee table: %s", err)
	}
	_, err = db.Exec(`INSERT INTO Department (id, name) VALUES (1, 'IT'), (2, 'Sales')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Department table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("SQL query returned error: %s", err)
	}

	type DepartmentHighestSalary struct {
		Department string
		Employee   string
		Salary     int
	}

	var results []DepartmentHighestSalary
	for rows.Next() {
		var result DepartmentHighestSalary
		if err := rows.Scan(&result.Department, &result.Employee, &result.Salary); err != nil {
			t.Fatalf("Failed to scan row: %s", err)
		}
		results = append(results, result)
	}

	expectedResults := []DepartmentHighestSalary{
		{"IT", "Jim", 90000},
		{"IT", "Max", 90000},
		{"Sales", "Henry", 80000},
	}

	if len(results) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(results))
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Department == results[j].Department {
			return results[i].Employee < results[j].Employee
		}
		return results[i].Department < results[j].Department
	})

	sort.Slice(expectedResults, func(i, j int) bool {
		if expectedResults[i].Department == expectedResults[j].Department {
			return expectedResults[i].Employee < expectedResults[j].Employee
		}
		return expectedResults[i].Department < expectedResults[j].Department
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

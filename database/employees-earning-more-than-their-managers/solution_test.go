package employees_earning_more_than_their_managers

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmployeesEarningMoreThanTheirManagers(t *testing.T) {
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
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Employee  (
  		Id int,
  		Name varchar(255),
  		Salary int,
  		ManagerId int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employee (Id, Name, Salary, ManagerId) VALUES 
		(1, 'Joe', 70000, 3), 
		(2, 'Henry', 80000, 4), 
		(3, 'Sam', 60000, NULL),
		(4, 'Max', 90000, NULL)
	`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Check the results of the query execution.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("Failed to select from table: %s", err)
	}

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		names = append(names, name)
	}

	expectedNames := []string{"Joe"}

	if len(names) != len(expectedNames) {
		t.Errorf("Expected %d names, got %d", len(expectedNames), len(names))
	}

	for i, name := range names {
		if name != expectedNames[i] {
			t.Errorf("Expected name at position %d to be %s, got %s", i, expectedNames[i], name)
		}
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

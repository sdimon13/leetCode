package replace_employee_id_with_the_unique_identifier

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestReplaceEmployeeIDWithUniqueIdentifier(t *testing.T) {
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

	// Drop the Employees and EmployeeUNI table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employees, EmployeeUNI`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employees and EmployeeUNI table.
	_, err = db.Exec(`CREATE TABLE Employees (
		id int PRIMARY KEY,
		name varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employees table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE EmployeeUNI (
		id int,
		unique_id int,
		PRIMARY KEY (id, unique_id)
	)`)

	if err != nil {
		t.Fatalf("Failed to create EmployeeUNI table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Employees (id, name) VALUES 
		(1, 'Alice'), 
		(7, 'Bob'),
		(11, 'Meir'),
		(90, 'Winston'),
		(3, 'Jonathan')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employees table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO EmployeeUNI (id, unique_id) VALUES 
		(3, 1), 
		(11, 2),
		(90, 3)`)

	if err != nil {
		t.Fatalf("Failed to insert data into EmployeeUNI table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]*int{
		"Alice":    nil,
		"Bob":      nil,
		"Meir":     func(i int) *int { return &i }(2),
		"Winston":  func(i int) *int { return &i }(3),
		"Jonathan": func(i int) *int { return &i }(1),
	}

	for rows.Next() {
		var unique_id *int
		var name string
		err := rows.Scan(&unique_id, &name)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedUniqueID, ok := expected[name]
		if !ok || (expectedUniqueID == nil && unique_id != nil) || (expectedUniqueID != nil && unique_id != nil && *expectedUniqueID != *unique_id) {
			t.Errorf("Unexpected result for name: %s. got unique_id: %v, want: %v", name, unique_id, expectedUniqueID)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

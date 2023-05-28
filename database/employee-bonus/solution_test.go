package employee_bonus

import (
	"database/sql"
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmployeeWithBonusLessThan1000(t *testing.T) {
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

	// Drop the Employee and Bonus table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employee, Bonus`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employee and Bonus table.
	_, err = db.Exec(`CREATE TABLE Employee (
		empId int PRIMARY KEY,
		name varchar(255),
		supervisor int,
		salary int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employee table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Bonus (
		empId int PRIMARY KEY,
		bonus int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Bonus table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Employee (empId, name, supervisor, salary) VALUES 
		(3, 'Brad', null, 4000),
		(1, 'John', 3, 1000),
		(2, 'Dan', 3, 2000),
		(4, 'Thomas', 3, 4000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employee table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Bonus (empId, bonus) VALUES 
		(2, 500),
		(4, 2000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Bonus table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]int{
		"Brad": 0,
		"John": 0,
		"Dan":  500,
	}

	for rows.Next() {
		var name string
		var bonus sql.NullInt64
		err := rows.Scan(&name, &bonus)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedBonus, ok := expected[name]
		if !ok || (bonus.Valid && int(bonus.Int64) != expectedBonus) {
			t.Errorf("Unexpected result for name: %s. got bonus: %d, want: %d", name, bonus.Int64, expectedBonus)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

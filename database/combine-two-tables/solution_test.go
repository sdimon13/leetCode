package combine_two_tables

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCombineTwoTables(t *testing.T) {
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

	// Drop the Person and Address tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Person, Address`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Person table.
	_, err = db.Exec(`CREATE TABLE Person (
		personId int PRIMARY KEY,
		lastName varchar(255),
		firstName varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Person table: %s", err)
	}

	// Create the Address table.
	_, err = db.Exec(`CREATE TABLE Address (
		addressId int PRIMARY KEY,
		personId int,
		city varchar(255),
		state varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Address table: %s", err)
	}

	// Populate the Person table with initial data.
	_, err = db.Exec(`INSERT INTO Person (personId, lastName, firstName) VALUES 
		(1, 'Wang', 'Allen'), 
		(2, 'Alice', 'Bob')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Person table: %s", err)
	}

	// Populate the Address table with initial data.
	_, err = db.Exec(`INSERT INTO Address (addressId, personId, city, state) VALUES 
		(1, 2, 'New York City', 'New York'), 
		(2, 3, 'Leetcode', 'California')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Address table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		firstName string
		lastName  string
		city      *string
		state     *string
	}{
		{"Allen", "Wang", nil, nil},
		{"Bob", "Alice", strPtr("New York City"), strPtr("New York")},
	}

	i := 0
	for rows.Next() {
		var firstName, lastName string
		var city, state *string
		err := rows.Scan(&firstName, &lastName, &city, &state)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if firstName != expected[i].firstName || lastName != expected[i].lastName ||
			(city == nil && expected[i].city != nil) || (city != nil && *city != *expected[i].city) ||
			(state == nil && expected[i].state != nil) || (state != nil && *state != *expected[i].state) {
			t.Errorf("Row %d - Expected (firstName: %s, lastName: %s, city: %v, state: %v), got (firstName: %s, lastName: %s, city: %v, state: %v)",
				i+1, expected[i].firstName, expected[i].lastName, expected[i].city, expected[i].state, firstName, lastName, city, state)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

func strPtr(s string) *string {
	return &s
}

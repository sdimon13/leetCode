package find_customer_referee

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFindCustomerReferee(t *testing.T) {
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

	// Create the Customer table.
	_, err = db.Exec(`CREATE TABLE Customer (
  		id int,
  		name varchar(255),
  		referee_id int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Customer (id, name, referee_id) VALUES 
		(1, 'Will', NULL), 
		(2, 'Jane', NULL), 
		(3, 'Alex', 2), 
		(4, 'Bill', NULL), 
		(5, 'Zack', 1), 
		(6, 'Mark', 2)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		names = append(names, name)
	}

	expectedNames := []string{"Will", "Jane", "Bill", "Zack"}

	if len(names) != len(expectedNames) {
		t.Errorf("Expected %d names, got %d", len(expectedNames), len(names))
	}

	sort.Strings(names)
	sort.Strings(expectedNames)

	for i, name := range names {
		if name != expectedNames[i] {
			t.Errorf("Expected name at position %d to be %s, got %s", i, expectedNames[i], name)
		}
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Customer`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

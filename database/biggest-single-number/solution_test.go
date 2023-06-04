package biggest_single_number

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBiggestSingleNumber(t *testing.T) {
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

	// Drop the MyNumbers table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS MyNumbers`)
	if err != nil {
		t.Fatalf("Failed to drop MyNumbers table: %s", err)
	}

	// Create the MyNumbers table.
	_, err = db.Exec(`CREATE TABLE MyNumbers (
		num int
	)`)
	if err != nil {
		t.Fatalf("Failed to create MyNumbers table: %s", err)
	}

	// Populate the MyNumbers table with initial data.
	_, err = db.Exec(`INSERT INTO MyNumbers (num) VALUES 
		(8), 
		(8), 
		(3), 
		(3),
		(1),
		(4),
		(5),
		(6)`)
	if err != nil {
		t.Fatalf("Failed to insert data into MyNumbers table: %s", err)
	}

	// Execute the SQL query from the file.
	row := db.QueryRow(string(query))

	var num *int
	err = row.Scan(&num)
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expectedNum := 6
	if num == nil || *num != expectedNum {
		t.Errorf("Expected %d, got %v", expectedNum, num)
	}

	// Now let's test the case where the result is null.

	// Clean up the MyNumbers table.
	_, err = db.Exec(`DELETE FROM MyNumbers`)
	if err != nil {
		t.Fatalf("Failed to clean up MyNumbers table: %s", err)
	}

	// Populate the MyNumbers table with new data.
	_, err = db.Exec(`INSERT INTO MyNumbers (num) VALUES 
		(8), 
		(8), 
		(7), 
		(7),
		(3),
		(3),
		(3)`)
	if err != nil {
		t.Fatalf("Failed to insert data into MyNumbers table: %s", err)
	}

	// Execute the SQL query from the file.
	row = db.QueryRow(string(query))

	num = nil
	err = row.Scan(&num)
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	if num != nil {
		t.Errorf("Expected nil, got %d", *num)
	}
}

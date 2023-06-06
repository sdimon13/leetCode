package last_person_to_fit_in_the_bus

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestLastPersonToBoardTheBus(t *testing.T) {
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

	// Drop the Queue table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Queue`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Queue table.
	_, err = db.Exec(`CREATE TABLE Queue (
		person_id int PRIMARY KEY,
		person_name varchar(255),
		weight int,
		turn int
	)`)
	if err != nil {
		t.Fatalf("Failed to create Queue table: %s", err)
	}

	// Populate the Queue table with initial data.
	_, err = db.Exec(`INSERT INTO Queue (person_id, person_name, weight, turn) VALUES 
		(5, 'Alice', 250, 1), 
		(4, 'Bob', 175, 5),
		(3, 'Alex', 350, 2),
		(6, 'John Cena', 400, 3),
		(1, 'Winston', 500, 6),
		(2, 'Marie', 200, 4)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Queue table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := "John Cena"

	for rows.Next() {
		var personName string
		err := rows.Scan(&personName)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if personName != expected {
			t.Errorf("Expected %s, got %s", expected, personName)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

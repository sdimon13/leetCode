package rising_temperature

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRisingTemperature(t *testing.T) {
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

	// Drop the Weather table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Weather`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Weather table.
	_, err = db.Exec(`CREATE TABLE Weather (
		id int PRIMARY KEY,
		recordDate date,
		temperature int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Weather table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Weather (id, recordDate, temperature) VALUES 
		(1, '2015-01-01', 10),
		(2, '2015-01-02', 25),
		(3, '2015-01-03', 20),
		(4, '2015-01-04', 30)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Weather table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]bool{
		2: true,
		4: true,
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		_, ok := expected[id]
		if !ok {
			t.Errorf("Unexpected result for id: %d. This id should not be in the result set.", id)
		} else {
			delete(expected, id)
		}
	}

	if len(expected) > 0 {
		for id := range expected {
			t.Errorf("Expected id %d was not in the result set.", id)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

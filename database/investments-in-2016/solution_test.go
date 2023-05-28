package investments_in_2016

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInvestmentsIn2016(t *testing.T) {
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

	// Drop the Insurance table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Insurance`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Insurance table.
	_, err = db.Exec(`CREATE TABLE Insurance (
		pid int PRIMARY KEY,
		tiv_2015 float,
		tiv_2016 float,
		lat float,
		lon float
	)`)
	if err != nil {
		t.Fatalf("Failed to create Insurance table: %s", err)
	}

	// Populate the Insurance table with initial data.
	_, err = db.Exec(`INSERT INTO Insurance (pid, tiv_2015, tiv_2016, lat, lon) VALUES 
		(1, 10, 5, 10, 10), 
		(2, 20, 20, 20, 20),
		(3, 10, 30, 20, 20),
		(4, 10, 40, 40, 40)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Insurance table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := 45.00

	if rows.Next() {
		var tiv_2016 float64
		err := rows.Scan(&tiv_2016)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if math.Abs(tiv_2016-expected) > 1e-9 {
			t.Errorf("Expected tiv_2016: %.2f, got %.2f", expected, tiv_2016)
		}
	} else {
		t.Errorf("No rows returned")
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

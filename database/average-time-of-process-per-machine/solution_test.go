package average_time_of_process_per_machine

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAverageTimeOfProcessPerMachine(t *testing.T) {
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

	// Drop the Activity table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Activity`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Activity table.
	_, err = db.Exec(`CREATE TABLE Activity (
		machine_id int,
		process_id int,
		activity_type ENUM('start', 'end'),
		timestamp float,
		PRIMARY KEY (machine_id, process_id, activity_type)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Activity table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Activity (machine_id, process_id, activity_type, timestamp) VALUES 
		(0, 0, 'start', 0.712),
		(0, 0, 'end', 1.520),
		(0, 1, 'start', 3.140),
		(0, 1, 'end', 4.120),
		(1, 0, 'start', 0.550),
		(1, 0, 'end', 1.550),
		(1, 1, 'start', 0.430),
		(1, 1, 'end', 1.420),
		(2, 0, 'start', 4.100),
		(2, 0, 'end', 4.512),
		(2, 1, 'start', 2.500),
		(2, 1, 'end', 5.000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Activity table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]float64{
		0: 0.894,
		1: 0.995,
		2: 1.456,
	}

	for rows.Next() {
		var machine_id int
		var processing_time float64
		err := rows.Scan(&machine_id, &processing_time)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedProcessingTime, ok := expected[machine_id]
		if !ok || math.Abs(processing_time-expectedProcessingTime) > 0.001 {
			t.Errorf("Unexpected result for machine_id: %d. got processing_time: %.3f, want: %.3f", machine_id, processing_time, expectedProcessingTime)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

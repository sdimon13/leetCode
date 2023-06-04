package game_play_analysis_iv

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGamePlayAnalysisIV(t *testing.T) {
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
		player_id int,
		device_id int,
		event_date date,
		games_played int,
		PRIMARY KEY (player_id, event_date)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Activity table: %s", err)
	}

	// Populate the Activity table with initial data.
	_, err = db.Exec(`INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES 
		(1, 2, '2016-03-01', 5), 
		(1, 2, '2016-03-02', 6), 
		(2, 3, '2017-06-25', 1), 
		(3, 1, '2016-03-02', 0), 
		(3, 4, '2018-07-03', 5)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Activity table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expectedFraction := 0.33

	if rows.Next() {
		var fraction float64
		err := rows.Scan(&fraction)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if math.Abs(fraction-expectedFraction) > 0.01 {
			t.Errorf("Expected fraction: %.2f, got: %.2f", expectedFraction, fraction)
		}
	} else {
		t.Errorf("No rows returned")
	}
}

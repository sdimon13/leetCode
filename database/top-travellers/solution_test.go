package top_travellers

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTopTravellers(t *testing.T) {
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

	// Drop the Users and Rides tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Users, Rides`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Users table.
	_, err = db.Exec(`CREATE TABLE Users (
		id int PRIMARY KEY,
		name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	// Create the Rides table.
	_, err = db.Exec(`CREATE TABLE Rides (
		id int PRIMARY KEY,
		user_id int,
		distance int
	)`)
	if err != nil {
		t.Fatalf("Failed to create Rides table: %s", err)
	}

	// Populate the Users table with initial data.
	_, err = db.Exec(`INSERT INTO Users (id, name) VALUES 
		(1, 'Alice'), 
		(2, 'Bob'),
		(3, 'Alex'),
		(4, 'Donald'),
		(7, 'Lee'),
		(13, 'Jonathan'),
		(19, 'Elvis')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	// Populate the Rides table with initial data.
	_, err = db.Exec(`INSERT INTO Rides (id, user_id, distance) VALUES 
		(1, 1, 120), 
		(2, 2, 317),
		(3, 3, 222),
		(4, 7, 100),
		(5, 13, 312),
		(6, 19, 50),
		(7, 7, 120),
		(8, 19, 400),
		(9, 7, 230)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Rides table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		name               string
		travelled_distance int
	}{
		{"Elvis", 450},
		{"Lee", 450},
		{"Bob", 317},
		{"Jonathan", 312},
		{"Alex", 222},
		{"Alice", 120},
		{"Donald", 0},
	}

	i := 0
	for rows.Next() {
		var name string
		var travelled_distance int
		err := rows.Scan(&name, &travelled_distance)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if name != expected[i].name || travelled_distance != expected[i].travelled_distance {
			t.Errorf("Row %d - Expected (name: %s, travelled_distance: %d), got (name: %s, travelled_distance: %d)",
				i+1, expected[i].name, expected[i].travelled_distance, name, travelled_distance)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

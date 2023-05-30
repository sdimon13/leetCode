package percentage_of_users_attended_a_contest

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPercentageOfUsersAttendedContest(t *testing.T) {
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

	// Drop the Users and Register tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Users, Register`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Users table.
	_, err = db.Exec(`CREATE TABLE Users (
		user_id int PRIMARY KEY,
		user_name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	// Create the Register table.
	_, err = db.Exec(`CREATE TABLE Register (
		contest_id int,
		user_id int,
		PRIMARY KEY (contest_id, user_id)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Register table: %s", err)
	}

	// Populate the Users table with initial data.
	_, err = db.Exec(`INSERT INTO Users (user_id, user_name) VALUES 
		(6, 'Alice'), 
		(2, 'Bob'), 
		(7, 'Alex')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	// Populate the Register table with initial data.
	_, err = db.Exec(`INSERT INTO Register (contest_id, user_id) VALUES 
		(215, 6), 
		(209, 2), 
		(208, 2), 
		(210, 6), 
		(208, 6), 
		(209, 7), 
		(209, 6), 
		(215, 7), 
		(208, 7), 
		(210, 2), 
		(207, 2), 
		(210, 7)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Register table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		contest_id int
		percentage float64
	}{
		{208, 100.0},
		{209, 100.0},
		{210, 100.0},
		{215, 66.67},
		{207, 33.33},
	}

	i := 0
	for rows.Next() {
		var contest_id int
		var percentage float64
		err := rows.Scan(&contest_id, &percentage)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if contest_id != expected[i].contest_id || math.Abs(percentage-expected[i].percentage) > 0.01 {
			t.Errorf("Row %d - Expected (contest_id: %d, percentage: %.2f), got (contest_id: %d, percentage: %.2f)",
				i+1, expected[i].contest_id, expected[i].percentage, contest_id, percentage)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

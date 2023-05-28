package confirmation_rate

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConfirmationRate(t *testing.T) {
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

	// Drop the Signups and Confirmations table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Signups, Confirmations`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Signups and Confirmations table.
	_, err = db.Exec(`CREATE TABLE Signups (
		user_id int PRIMARY KEY,
		time_stamp datetime
	)`)

	if err != nil {
		t.Fatalf("Failed to create Signups table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Confirmations (
		user_id int,
		time_stamp datetime,
		action ENUM('confirmed', 'timeout'),
		PRIMARY KEY(user_id, time_stamp)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Confirmations table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Signups (user_id, time_stamp) VALUES 
		(3, '2020-03-21 10:16:13'), 
		(7, '2020-01-04 13:57:59'),
		(2, '2020-07-29 23:09:44'),
		(6, '2020-12-09 10:39:37')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Signups table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Confirmations (user_id, time_stamp, action) VALUES 
		(3, '2021-01-06 03:30:46', 'timeout'), 
		(3, '2021-07-14 14:00:00', 'timeout'),
		(7, '2021-06-12 11:57:29', 'confirmed'),
		(7, '2021-06-13 12:58:28', 'confirmed'),
		(7, '2021-06-14 13:59:27', 'confirmed'),
		(2, '2021-01-22 00:00:00', 'confirmed'),
		(2, '2021-02-28 23:59:59', 'timeout')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Confirmations table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]float64{
		6: 0.00,
		3: 0.00,
		7: 1.00,
		2: 0.50,
	}

	for rows.Next() {
		var user_id int
		var confirmation_rate float64
		err := rows.Scan(&user_id, &confirmation_rate)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedRate, ok := expected[user_id]
		if !ok || math.Abs(confirmation_rate-expectedRate) > 1e-2 {
			t.Errorf("Unexpected result for user_id: %d. got confirmation_rate: %.2f, want: %.2f", user_id, confirmation_rate, expectedRate)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

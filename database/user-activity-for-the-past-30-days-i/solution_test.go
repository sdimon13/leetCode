package user_activity_for_the_past_30_days_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestActiveUsers(t *testing.T) {
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
  		user_id int,
  		session_id int,
		activity_date date,
		activity_type enum('open_session', 'end_session', 'scroll_down', 'send_message')
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Activity (user_id, session_id, activity_date, activity_type) VALUES 
		(1, 1, '2019-07-20', 'open_session'), 
		(1, 1, '2019-07-20', 'scroll_down'),
		(1, 1, '2019-07-20', 'end_session'),
		(2, 4, '2019-07-20', 'open_session'),
		(2, 4, '2019-07-21', 'send_message'),
		(2, 4, '2019-07-21', 'end_session'),
		(3, 2, '2019-07-21', 'open_session'),
		(3, 2, '2019-07-21', 'send_message'),
		(3, 2, '2019-07-21', 'end_session'),
		(4, 3, '2019-06-25', 'open_session'),
		(4, 3, '2019-06-25', 'end_session')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		day          string
		active_users int
	}{
		{"2019-07-20", 2},
		{"2019-07-21", 2},
	}

	i := 0
	for rows.Next() {
		var day string
		var active_users int
		err := rows.Scan(&day, &active_users)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if day != expected[i].day || active_users != expected[i].active_users {
			t.Errorf("Expected row %d to be %v, got %s %d", i+1, expected[i], day, active_users)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

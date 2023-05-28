package friend_requests_ii_who_has_the_most_friends

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMostFriends(t *testing.T) {
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

	// Drop the RequestAccepted table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS RequestAccepted`)
	if err != nil {
		t.Fatalf("Failed to drop RequestAccepted table: %s", err)
	}

	// Create the RequestAccepted table.
	_, err = db.Exec(`CREATE TABLE RequestAccepted (
		requester_id int,
		accepter_id int,
		accept_date date,
		PRIMARY KEY (requester_id, accepter_id)
	)`)
	if err != nil {
		t.Fatalf("Failed to create RequestAccepted table: %s", err)
	}

	// Populate the RequestAccepted table with initial data.
	_, err = db.Exec(`INSERT INTO RequestAccepted (requester_id, accepter_id, accept_date) VALUES 
		(1, 2, '2016/06/03'), 
		(1, 3, '2016/06/08'),
		(2, 3, '2016/06/08'),
		(3, 4, '2016/06/09')`)
	if err != nil {
		t.Fatalf("Failed to insert data into RequestAccepted table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		id  int
		num int
	}{
		{3, 3},
	}

	i := 0
	for rows.Next() {
		var id, num int
		err := rows.Scan(&id, &num)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id || num != expected[i].num {
			t.Errorf("Row %d - Expected (id: %d, num: %d), got (id: %d, num: %d)",
				i+1, expected[i].id, expected[i].num, id, num)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

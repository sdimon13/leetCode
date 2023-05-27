package find_followers_count

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFindFollowersCount(t *testing.T) {
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

	// Drop the Followers table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Followers`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Followers table.
	_, err = db.Exec(`CREATE TABLE Followers (
		user_id int,
		follower_id int,
		PRIMARY KEY (user_id, follower_id)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Followers table: %s", err)
	}

	// Populate the Followers table with initial data.
	_, err = db.Exec(`INSERT INTO Followers (user_id, follower_id) VALUES 
		(0, 1), 
		(1, 0),
		(2, 0),
		(2, 1)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Followers table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		user_id         int
		followers_count int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
	}

	i := 0
	for rows.Next() {
		var user_id, followers_count int
		err := rows.Scan(&user_id, &followers_count)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if user_id != expected[i].user_id || followers_count != expected[i].followers_count {
			t.Errorf("Row %d - Expected (user_id: %d, followers_count: %d), got (user_id: %d, followers_count: %d)",
				i+1, expected[i].user_id, expected[i].followers_count, user_id, followers_count)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

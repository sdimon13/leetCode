package invalid_tweets

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInvalidTweets(t *testing.T) {
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

	// Drop the Tweets table if it exists.
	_, err = db.Exec("DROP TABLE IF EXISTS Tweets")
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Tweets table.
	_, err = db.Exec(`CREATE TABLE Tweets (
		tweet_id int PRIMARY KEY,
		content varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Tweets table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Tweets (tweet_id, content) VALUES 
		(1, 'Vote for Biden'),
		(2, 'Let us make America great again!')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Tweets table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []int{2}

	for rows.Next() {
		var tweetID int
		err := rows.Scan(&tweetID)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		found := false
		for _, id := range expected {
			if tweetID == id {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Unexpected result: got tweet_id %d, want one of %v", tweetID, expected)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

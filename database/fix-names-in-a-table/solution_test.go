package duplicate_emails

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFixNamesInTable(t *testing.T) {
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

	// Create the Users table.
	_, err = db.Exec(`CREATE TABLE Users (
  		user_id int,
  		name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Users (user_id, name) VALUES 
		(1, 'aLice'), 
		(2, 'bOB')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		user_id int
		name    string
	}{
		{1, "Alice"},
		{2, "Bob"},
	}

	i := 0
	for rows.Next() {
		var user_id int
		var name string
		err := rows.Scan(&user_id, &name)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if user_id != expected[i].user_id || name != expected[i].name {
			t.Errorf("Expected row %d to be %v, got %d %s", i+1, expected[i], user_id, name)
		}
		i++
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Users`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

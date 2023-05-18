package delete_duplicate_emails

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDeleteDuplicateEmails(t *testing.T) {
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

	// Create the Person table.
	_, err = db.Exec(`CREATE TABLE Person (
  		Id int,
  		Email varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Person (Id, Email) VALUES (1, 'john@example.com'), (2, 'bob@example.com'), (3, 'john@example.com')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	_, err = db.Exec(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	// Check the results of the query execution.
	rows, err := db.Query("SELECT Email FROM Person ORDER BY Email")
	if err != nil {
		t.Errorf("Failed to select from table: %s", err)
	}

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		emails = append(emails, email)
	}

	expectedEmails := []string{"bob@example.com", "john@example.com"}

	if len(emails) != len(expectedEmails) {
		t.Errorf("Expected %d emails, got %d", len(expectedEmails), len(emails))
	}

	for i, email := range emails {
		if email != expectedEmails[i] {
			t.Errorf("Expected email at position %d to be %s, got %s", i, expectedEmails[i], email)
		}
	}

	if rows.Next() {
		t.Errorf("Duplicate emails found")
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Person`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

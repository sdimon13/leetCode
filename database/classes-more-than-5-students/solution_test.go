package classes_more_than_5_students

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestClassesWithAtLeastFiveStudents(t *testing.T) {
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

	// Drop the Courses table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Courses`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Courses table.
	_, err = db.Exec(`CREATE TABLE Courses (
		student varchar(255) PRIMARY KEY,
		class varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Courses table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Courses (student, class) VALUES 
		('A', 'Math'), 
		('B', 'English'),
		('C', 'Math'),
		('D', 'Biology'),
		('E', 'Math'),
		('F', 'Computer'),
		('G', 'Math'),
		('H', 'Math'),
		('I', 'Math')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Courses table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]bool{
		"Math": true,
	}

	for rows.Next() {
		var class string
		err := rows.Scan(&class)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		_, ok := expected[class]
		if !ok {
			t.Errorf("Unexpected result for class: %s.", class)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package number_of_unique_subjects_taught_by_each_teacher

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNumberOfUniqueSubjectsTaughtByEachTeacher(t *testing.T) {
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

	// Drop the Teacher table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Teacher`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Teacher table.
	_, err = db.Exec(`CREATE TABLE Teacher (
		teacher_id int,
		subject_id int,
		dept_id int,
		PRIMARY KEY (subject_id, dept_id)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Teacher table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Teacher (teacher_id, subject_id, dept_id) VALUES 
		(1, 2, 3), 
		(1, 2, 4),
		(1, 3, 3),
		(2, 1, 1),
		(2, 2, 1),
		(2, 3, 1),
		(2, 4, 1)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Teacher table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]int{
		1: 2,
		2: 4,
	}

	for rows.Next() {
		var teacher_id int
		var count_subjects int
		err := rows.Scan(&teacher_id, &count_subjects)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedCount, ok := expected[teacher_id]
		if !ok || count_subjects != expectedCount {
			t.Errorf("Unexpected result for teacher_id: %d. got count_subjects: %d, want: %d", teacher_id, count_subjects, expectedCount)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package students_and_examinations

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestStudentsAndExaminations(t *testing.T) {
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

	// Drop the Students, Subjects and Examinations tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Students, Subjects, Examinations`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Students table.
	_, err = db.Exec(`CREATE TABLE Students (
		student_id int PRIMARY KEY,
		student_name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Students table: %s", err)
	}

	// Create the Subjects table.
	_, err = db.Exec(`CREATE TABLE Subjects (
		subject_name varchar(255) PRIMARY KEY
	)`)
	if err != nil {
		t.Fatalf("Failed to create Subjects table: %s", err)
	}

	// Create the Examinations table.
	_, err = db.Exec(`CREATE TABLE Examinations (
		student_id int,
		subject_name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Examinations table: %s", err)
	}

	// Populate the Students table with initial data.
	_, err = db.Exec(`INSERT INTO Students (student_id, student_name) VALUES 
		(1, 'Alice'), 
		(2, 'Bob'),
		(13, 'John'),
		(6, 'Alex')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Students table: %s", err)
	}

	// Populate the Subjects table with initial data.
	_, err = db.Exec(`INSERT INTO Subjects (subject_name) VALUES 
		('Math'), 
		('Physics'),
		('Programming')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Subjects table: %s", err)
	}

	// Populate the Examinations table with initial data.
	_, err = db.Exec(`INSERT INTO Examinations (student_id, subject_name) VALUES 
		(1, 'Math'), 
		(1, 'Physics'),
		(1, 'Programming'),
		(2, 'Programming'),
		(1, 'Physics'),
		(1, 'Math'),
		(13, 'Math'),
		(13, 'Programming'),
		(13, 'Physics'),
		(2, 'Math'),
		(1, 'Math')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Examinations table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		studentId     int
		studentName   string
		subjectName   string
		attendedExams int
	}{
		{1, "Alice", "Math", 3},
		{1, "Alice", "Physics", 2},
		{1, "Alice", "Programming", 1},
		{2, "Bob", "Math", 1},
		{2, "Bob", "Physics", 0},
		{2, "Bob", "Programming", 1},
		{6, "Alex", "Math", 0},
		{6, "Alex", "Physics", 0},
		{6, "Alex", "Programming", 0},
		{13, "John", "Math", 1},
		{13, "John", "Physics", 1},
		{13, "John", "Programming", 1},
	}

	i := 0
	for rows.Next() {
		var studentId, attendedExams int
		var studentName, subjectName string
		err := rows.Scan(&studentId, &studentName, &subjectName, &attendedExams)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if studentId != expected[i].studentId || studentName != expected[i].studentName ||
			subjectName != expected[i].subjectName || attendedExams != expected[i].attendedExams {
			t.Errorf("Row %d - Expected (studentId: %d, studentName: %s, subjectName: %s, attendedExams: %d), got (studentId: %d, studentName: %s, subjectName: %s, attendedExams: %d)",
				i+1, expected[i].studentId, expected[i].studentName, expected[i].subjectName, expected[i].attendedExams, studentId, studentName, subjectName, attendedExams)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

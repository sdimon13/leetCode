package find_total_time_spent_by_each_employee

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTotalTimeSpentByEachEmployee(t *testing.T) {
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

	// Drop the Employees table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Employees`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Employees table.
	_, err = db.Exec(`CREATE TABLE Employees (
		emp_id int,
		event_day date,
		in_time int,
		out_time int,
		PRIMARY KEY(emp_id, event_day, in_time)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Employees table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employees (emp_id, event_day, in_time, out_time) VALUES 
		(1, '2020-11-28', 4, 32),
		(1, '2020-11-28', 55, 200),
		(1, '2020-12-03', 1, 42),
		(2, '2020-11-28', 3, 33),
		(2, '2020-12-09', 47, 74)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Employees table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]map[int]int{
		"2020-11-28": {1: 173, 2: 30},
		"2020-12-03": {1: 41},
		"2020-12-09": {2: 27},
	}

	for rows.Next() {
		var day string
		var emp_id int
		var total_time int
		err := rows.Scan(&day, &emp_id, &total_time)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedTime, ok := expected[day][emp_id]
		if !ok || total_time != expectedTime {
			t.Errorf("Unexpected result for day: %s, emp_id: %d. got total_time: %d, want: %d", day, emp_id, total_time, expectedTime)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

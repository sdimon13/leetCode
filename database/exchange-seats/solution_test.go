package exchange_seats

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExchangeSeats(t *testing.T) {
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

	// Drop the Seat table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Seat`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Seat table.
	_, err = db.Exec(`CREATE TABLE Seat (
		id int PRIMARY KEY,
		student varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Seat table: %s", err)
	}

	// Populate the Seat table with initial data.
	_, err = db.Exec(`INSERT INTO Seat (id, student) VALUES 
		(1, 'Abbot'), 
		(2, 'Doris'),
		(3, 'Emerson'),
		(4, 'Green'),
		(5, 'Jeames')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Seat table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		id      int
		student string
	}{
		{1, "Doris"},
		{2, "Abbot"},
		{3, "Green"},
		{4, "Emerson"},
		{5, "Jeames"},
	}

	i := 0
	for rows.Next() {
		var id int
		var student string
		err := rows.Scan(&id, &student)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id || student != expected[i].student {
			t.Errorf("Row %d - Expected (id: %d, student: %s), got (id: %d, student: %s)",
				i+1, expected[i].id, expected[i].student, id, student)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

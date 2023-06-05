package consecutive_numbers

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConsecutiveNums(t *testing.T) {
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

	// Drop the Logs table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Logs`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Logs table.
	_, err = db.Exec(`CREATE TABLE Logs (
		id int PRIMARY KEY AUTO_INCREMENT,
		num varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Logs table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Logs (num) VALUES 
		("1"), 
		("1"),
		("1"),
		("2"),
		("1"),
		("2"),
		("2")`)

	if err != nil {
		t.Fatalf("Failed to insert data into Logs table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]bool{
		"1": true,
	}

	for rows.Next() {
		var num string
		err := rows.Scan(&num)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		_, ok := expected[num]
		if !ok {
			t.Errorf("Unexpected result for num: %s", num)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

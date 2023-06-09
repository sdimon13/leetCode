package triangle_judgement

import (
	"fmt"
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTriangleJudgement(t *testing.T) {
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

	// Drop the Triangle table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Triangle`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Triangle table.
	_, err = db.Exec(`CREATE TABLE Triangle (
		x int,
		y int,
		z int,
		PRIMARY KEY (x, y, z)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Triangle table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Triangle (x, y, z) VALUES 
		(13, 15, 30), 
		(10, 20, 15)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Triangle table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]string{
		"13,15,30": "No",
		"10,20,15": "Yes",
	}

	for rows.Next() {
		var x, y, z int
		var triangle string
		err := rows.Scan(&x, &y, &z, &triangle)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedTriangle, ok := expected[fmt.Sprintf("%d,%d,%d", x, y, z)]
		if !ok || triangle != expectedTriangle {
			t.Errorf("Unexpected result for triangle: %d,%d,%d. got: %s, want: %s", x, y, z, triangle, expectedTriangle)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

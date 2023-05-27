package tree_node

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTreeNode(t *testing.T) {
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

	// Drop the Tree table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Tree`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Tree table.
	_, err = db.Exec(`CREATE TABLE Tree (
		id int PRIMARY KEY,
		p_id int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Tree (id, p_id) VALUES 
		(1, NULL), 
		(2, 1),
		(3, 1),
		(4, 2),
		(5, 2)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		id       int
		nodeType string
	}{
		{1, "Root"},
		{2, "Inner"},
		{3, "Leaf"},
		{4, "Leaf"},
		{5, "Leaf"},
	}

	i := 0
	for rows.Next() {
		var id int
		var nodeType string
		err := rows.Scan(&id, &nodeType)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id || nodeType != expected[i].nodeType {
			t.Errorf("Row %d - Expected (id: %d, type: %s), got (id: %d, type: %s)",
				i+1, expected[i].id, expected[i].nodeType, id, nodeType)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

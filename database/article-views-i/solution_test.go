package article_views_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestArticleViews(t *testing.T) {
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

	// Drop the Views table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Views`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Views table.
	_, err = db.Exec(`CREATE TABLE Views (
		article_id int,
		author_id int,
		viewer_id int,
		view_date date
	)`)
	if err != nil {
		t.Fatalf("Failed to create Views table: %s", err)
	}

	// Populate the Views table with initial data.
	_, err = db.Exec(`INSERT INTO Views (article_id, author_id, viewer_id, view_date) VALUES 
		(1, 3, 5, '2019-08-01'),
		(1, 3, 6, '2019-08-02'),
		(2, 7, 7, '2019-08-01'),
		(2, 7, 6, '2019-08-02'),
		(4, 7, 1, '2019-07-22'),
		(3, 4, 4, '2019-07-21'),
		(3, 4, 4, '2019-07-21')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Views table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		id int
	}{
		{4},
		{7},
	}

	i := 0
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id {
			t.Errorf("Row %d - Expected (id: %d), got (id: %d)",
				i+1, expected[i].id, id)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package not_boring_movies

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNotBoringMovies(t *testing.T) {
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

	// Drop the Cinema table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Cinema`)
	if err != nil {
		t.Fatalf("Failed to drop Cinema table: %s", err)
	}

	// Create the Cinema table.
	_, err = db.Exec(`CREATE TABLE Cinema (
		id int PRIMARY KEY,
		movie varchar(255),
		description varchar(255),
		rating float
	)`)
	if err != nil {
		t.Fatalf("Failed to create Cinema table: %s", err)
	}

	// Populate the Cinema table with initial data.
	_, err = db.Exec(`INSERT INTO Cinema (id, movie, description, rating) VALUES 
		(1, 'War', 'great 3D', 8.9), 
		(2, 'Science', 'fiction', 8.5),
		(3, 'irish', 'boring', 6.2),
		(4, 'Ice song', 'Fantacy', 8.6),
		(5, 'House card', 'Interesting', 9.1)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Cinema table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		id          int
		movie       string
		description string
		rating      float64
	}{
		{5, "House card", "Interesting", 9.1},
		{1, "War", "great 3D", 8.9},
	}

	i := 0
	for rows.Next() {
		var id int
		var movie, description string
		var rating float64
		err := rows.Scan(&id, &movie, &description, &rating)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id || movie != expected[i].movie || description != expected[i].description || rating != expected[i].rating {
			t.Errorf("Row %d - Expected (id: %d, movie: %s, description: %s, rating: %f), got (id: %d, movie: %s, description: %s, rating: %f)",
				i+1, expected[i].id, expected[i].movie, expected[i].description, expected[i].rating, id, movie, description, rating)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

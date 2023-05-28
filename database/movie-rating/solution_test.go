package movie_rating

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMovieRating(t *testing.T) {
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

	// Drop the Movies, Users, and MovieRating tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Movies, Users, MovieRating`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Movies table.
	_, err = db.Exec(`CREATE TABLE Movies (
		movie_id int PRIMARY KEY,
		title varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Movies table: %s", err)
	}

	// Create the Users table.
	_, err = db.Exec(`CREATE TABLE Users (
		user_id int PRIMARY KEY,
		name varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	// Create the MovieRating table.
	_, err = db.Exec(`CREATE TABLE MovieRating (
		movie_id int,
		user_id int,
		rating int,
		created_at date,
		PRIMARY KEY (movie_id, user_id)
	)`)
	if err != nil {
		t.Fatalf("Failed to create MovieRating table: %s", err)
	}

	// Populate the Movies table with initial data.
	_, err = db.Exec(`INSERT INTO Movies (movie_id, title) VALUES 
		(1, 'Avengers'), 
		(2, 'Frozen 2'),
		(3, 'Joker')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Movies table: %s", err)
	}

	// Populate the Users table with initial data.
	_, err = db.Exec(`INSERT INTO Users (user_id, name) VALUES 
		(1, 'Daniel'), 
		(2, 'Monica'),
		(3, 'Maria'),
		(4, 'James')`)
	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	// Populate the MovieRating table with initial data.
	_, err = db.Exec(`INSERT INTO MovieRating (movie_id, user_id, rating, created_at) VALUES 
		(1, 1, 3, '2020-01-12'), 
		(1, 2, 4, '2020-02-11'),
		(1, 3, 2, '2020-02-12'),
		(1, 4, 1, '2020-01-01'),
		(2, 1, 5, '2020-02-17'),
		(2, 2, 2, '2020-02-01'),
		(2, 3, 2, '2020-03-01'),
		(3, 1, 3, '2020-02-22'),
		(3, 2, 4, '2020-02-25')`)
	if err != nil {
		t.Fatalf("Failed to insert data into MovieRating table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		results string
	}{
		{"Daniel"},
		{"Frozen 2"},
	}

	i := 0
	for rows.Next() {
		var results string
		err := rows.Scan(&results)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if results != expected[i].results {
			t.Errorf("Row %d - Expected (results: %s), got (results: %s)",
				i+1, expected[i].results, results)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

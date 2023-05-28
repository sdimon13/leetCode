package actors_and_directors_who_cooperated_at_least_three_times

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestActorDirectorCooperation(t *testing.T) {
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

	// Drop the ActorDirector table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS ActorDirector`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the ActorDirector table.
	_, err = db.Exec(`CREATE TABLE ActorDirector (
		actor_id int,
		director_id int,
		timestamp int PRIMARY KEY
	)`)

	if err != nil {
		t.Fatalf("Failed to create ActorDirector table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO ActorDirector (actor_id, director_id, timestamp) VALUES 
		(1, 1, 0),
		(1, 1, 1),
		(1, 1, 2),
		(1, 2, 3),
		(1, 2, 4),
		(2, 1, 5),
		(2, 1, 6)`)

	if err != nil {
		t.Fatalf("Failed to insert data into ActorDirector table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]int{
		1: 1,
	}

	for rows.Next() {
		var actor_id int
		var director_id int
		err := rows.Scan(&actor_id, &director_id)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedDirector, ok := expected[actor_id]
		if !ok || director_id != expectedDirector {
			t.Errorf("Unexpected result for actor_id: %d. got director_id: %d, want: %d", actor_id, director_id, expectedDirector)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

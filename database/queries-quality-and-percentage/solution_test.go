package queries_quality_and_percentage

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestQueryQualityAndPercentage(t *testing.T) {
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

	// Drop the Queries table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Queries`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Queries table.
	_, err = db.Exec(`CREATE TABLE Queries (
		query_name varchar(255),
		result varchar(255),
		position int,
		rating int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Queries table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Queries (query_name, result, position, rating) VALUES 
		('Dog', 'Golden Retriever', 1, 5),
		('Dog', 'German Shepherd', 2, 5),
		('Dog', 'Mule', 200, 1),
		('Cat', 'Shirazi', 5, 2),
		('Cat', 'Siamese', 3, 3),
		('Cat', 'Sphynx', 7, 4)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Queries table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]struct {
		Quality             float64
		PoorQueryPercentage float64
	}{
		"Dog": {Quality: 2.50, PoorQueryPercentage: 33.33},
		"Cat": {Quality: 0.66, PoorQueryPercentage: 33.33},
	}

	for rows.Next() {
		var queryName string
		var quality, poorQueryPercentage float64
		err := rows.Scan(&queryName, &quality, &poorQueryPercentage)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedResult, ok := expected[queryName]
		if !ok || math.Abs(quality-expectedResult.Quality) > 0.01 || math.Abs(poorQueryPercentage-expectedResult.PoorQueryPercentage) > 0.01 {
			t.Errorf("Unexpected result for query_name: %s. got quality: %.2f, poor_query_percentage: %.2f, want quality: %.2f, poor_query_percentage: %.2f", queryName, quality, poorQueryPercentage, expectedResult.Quality, expectedResult.PoorQueryPercentage)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

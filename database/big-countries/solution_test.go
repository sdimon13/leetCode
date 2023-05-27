package big_countries

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBigCountries(t *testing.T) {
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

	// Create the World table.
	_, err = db.Exec(`CREATE TABLE World (
  		name varchar(255),
  		continent varchar(255),
  		area int,
  		population int,
  		gdp bigint
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO World (name, continent, area, population, gdp) VALUES 
		('Afghanistan', 'Asia', 652230, 25500100, 20343000000), 
		('Albania', 'Europe', 28748, 2831741, 12960000000), 
		('Algeria', 'Africa', 2381741, 37100000, 188681000000), 
		('Andorra', 'Europe', 468, 78115, 3712000000), 
		('Angola', 'Africa', 1246700, 20609294, 100990000000)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	var countries []struct {
		name       string
		population int
		area       int
	}
	for rows.Next() {
		var c struct {
			name       string
			population int
			area       int
		}
		if err := rows.Scan(&c.name, &c.population, &c.area); err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}
		countries = append(countries, c)
	}

	expectedCountries := []struct {
		name       string
		population int
		area       int
	}{
		{"Afghanistan", 25500100, 652230},
		{"Algeria", 37100000, 2381741},
	}

	if len(countries) != len(expectedCountries) {
		t.Errorf("Expected %d countries, got %d", len(expectedCountries), len(countries))
	}

	for i, country := range countries {
		if country.name != expectedCountries[i].name ||
			country.population != expectedCountries[i].population ||
			country.area != expectedCountries[i].area {
			t.Errorf("Expected country at position %d to be %+v, got %+v", i, expectedCountries[i], country)
		}
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE World`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

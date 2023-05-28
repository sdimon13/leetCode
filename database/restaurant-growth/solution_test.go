package restaurant_growth

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMovingAverage(t *testing.T) {
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

	// Drop the Customer table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Customer`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Customer table.
	_, err = db.Exec(`CREATE TABLE Customer (
		customer_id int,
		name varchar(255),
		visited_on date,
		amount int,
		PRIMARY KEY (customer_id, visited_on)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Customer table: %s", err)
	}

	// Populate the Customer table with initial data.
	_, err = db.Exec(`INSERT INTO Customer (customer_id, name, visited_on, amount) VALUES 
		(1, 'Jhon', '2019-01-01', 100), 
		(2, 'Daniel', '2019-01-02', 110),
		(3, 'Jade', '2019-01-03', 120),
		(4, 'Khaled', '2019-01-04', 130),
		(5, 'Winston', '2019-01-05', 110), 
		(6, 'Elvis', '2019-01-06', 140), 
		(7, 'Anna', '2019-01-07', 150),
		(8, 'Maria', '2019-01-08', 80),
		(9, 'Jaze', '2019-01-09', 110), 
		(1, 'Jhon', '2019-01-10', 130), 
		(3, 'Jade', '2019-01-10', 150)`)
	if err != nil {
		t.Fatalf("Failed to insert data into Customer table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		visitedOn     string
		averageAmount float64
	}{
		{"2019-01-07", 122.86},
		{"2019-01-08", 120},
		{"2019-01-09", 120},
		{"2019-01-10", 142.86},
	}

	i := 0
	for rows.Next() {
		var visitedOn string
		var averageAmount float64
		err := rows.Scan(&visitedOn, &averageAmount)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if visitedOn != expected[i].visitedOn || averageAmount != expected[i].averageAmount {
			t.Errorf("Row %d - Expected (visitedOn: %s, averageAmount: %.2f), got (visitedOn: %s, averageAmount: %.2f)",
				i+1, expected[i].visitedOn, expected[i].averageAmount, visitedOn, averageAmount)
		}
		i++
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

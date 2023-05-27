package group_sold_products_by_the_date

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGroupSoldProductsByDate(t *testing.T) {
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

	// Drop the Activities table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Activities`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Activities table.
	_, err = db.Exec(`CREATE TABLE Activities (
  		sell_date date,
  		product varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Activities (sell_date, product) VALUES 
		('2020-05-30', 'Headphone'), 
		('2020-06-01', 'Pencil'),
		('2020-06-02', 'Mask'),
		('2020-05-30', 'Basketball'),
		('2020-06-01', 'Bible'),
		('2020-06-02', 'Mask'),
		('2020-05-30', 'T-Shirt')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		sell_date string
		num_sold  int
		products  string
	}{
		{"2020-05-30", 3, "Basketball,Headphone,T-Shirt"},
		{"2020-06-01", 2, "Bible,Pencil"},
		{"2020-06-02", 1, "Mask"},
	}

	i := 0
	for rows.Next() {
		var sell_date string
		var num_sold int
		var products string
		err := rows.Scan(&sell_date, &num_sold, &products)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if sell_date != expected[i].sell_date || num_sold != expected[i].num_sold || products != expected[i].products {
			t.Errorf("Expected row %d to be %v, got %s %d %s", i+1, expected[i], sell_date, num_sold, products)
		}
		i++
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

package average_selling_price

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"math"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAverageSellingPrice(t *testing.T) {
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

	// Drop the Prices and UnitsSold tables if they exist.
	_, err = db.Exec(`DROP TABLE IF EXISTS Prices, UnitsSold`)
	if err != nil {
		t.Fatalf("Failed to drop tables: %s", err)
	}

	// Create the Prices and UnitsSold tables.
	_, err = db.Exec(`CREATE TABLE Prices (
		product_id int,
		start_date date,
		end_date date,
		price int,
		PRIMARY KEY (product_id, start_date, end_date)
	)`)
	if err != nil {
		t.Fatalf("Failed to create Prices table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE UnitsSold (
		product_id int,
		purchase_date date,
		units int
	)`)
	if err != nil {
		t.Fatalf("Failed to create UnitsSold table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Prices (product_id, start_date, end_date, price) VALUES 
		(1, '2019-02-17', '2019-02-28', 5), 
		(1, '2019-03-01', '2019-03-22', 20),
		(2, '2019-02-01', '2019-02-20', 15),
		(2, '2019-02-21', '2019-03-31', 30)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Prices table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO UnitsSold (product_id, purchase_date, units) VALUES 
		(1, '2019-02-25', 100), 
		(1, '2019-03-01', 15),
		(2, '2019-02-10', 200),
		(2, '2019-03-22', 30)`)

	if err != nil {
		t.Fatalf("Failed to insert data into UnitsSold table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]float64{
		1: 6.96,
		2: 16.96,
	}

	for rows.Next() {
		var product_id int
		var average_price float64
		err := rows.Scan(&product_id, &average_price)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedPrice, ok := expected[product_id]
		if !ok || math.Abs(average_price-expectedPrice) > 0.01 {
			t.Errorf("Unexpected result for product_id: %d. got average_price: %.2f, want: %.2f", product_id, average_price, expectedPrice)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

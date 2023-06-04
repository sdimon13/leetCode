package immediate_food_delivery_ii

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestImmediateDeliveryPercentage(t *testing.T) {
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

	// Drop the Delivery table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Delivery`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Delivery table.
	_, err = db.Exec(`CREATE TABLE Delivery (
		delivery_id int PRIMARY KEY,
		customer_id int,
		order_date date,
		customer_pref_delivery_date date
	)`)

	if err != nil {
		t.Fatalf("Failed to create Delivery table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Delivery (delivery_id, customer_id, order_date, customer_pref_delivery_date) VALUES 
		(1, 1, '2019-08-01', '2019-08-02'), 
		(2, 2, '2019-08-02', '2019-08-02'),
		(3, 1, '2019-08-11', '2019-08-12'),
		(4, 3, '2019-08-24', '2019-08-24'),
		(5, 3, '2019-08-21', '2019-08-22'),
		(6, 2, '2019-08-11', '2019-08-13'),
		(7, 4, '2019-08-09', '2019-08-09')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Delivery table: %s", err)
	}

	// Execute the SQL query from the file.
	row := db.QueryRow(string(query))

	var immediatePercentage float64
	err = row.Scan(&immediatePercentage)
	if err != nil {
		t.Fatalf("Failed to scan row: %s", err)
	}

	expectedImmediatePercentage := 50.00
	if immediatePercentage != expectedImmediatePercentage {
		t.Errorf("Unexpected immediate_percentage: got %.2f, want %.2f", immediatePercentage, expectedImmediatePercentage)
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

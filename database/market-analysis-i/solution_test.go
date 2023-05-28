package market_analysis_i

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserOrdersIn2019(t *testing.T) {
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

	// Drop the Users, Orders, and Items table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Users, Orders, Items`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Users, Orders, and Items table.
	_, err = db.Exec(`CREATE TABLE Users (
		user_id int PRIMARY KEY,
		join_date date,
		favorite_brand varchar(50)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Orders (
		order_id int PRIMARY KEY,
		order_date date,
		item_id int,
		buyer_id int,
		seller_id int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Orders table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Items (
		item_id int PRIMARY KEY,
		item_brand varchar(50)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Items table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Users (user_id, join_date, favorite_brand) VALUES 
		(1, '2018-01-01', 'Lenovo'), 
		(2, '2018-02-09', 'Samsung'),
		(3, '2018-01-19', 'LG'),
		(4, '2018-05-21', 'HP')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Orders (order_id, order_date, item_id, buyer_id, seller_id) VALUES 
		(1, '2019-08-01', 4, 1, 2),
		(2, '2018-08-02', 2, 1, 3),
		(3, '2019-08-03', 3, 2, 3),
		(4, '2018-08-04', 1, 4, 2),
		(5, '2018-08-04', 1, 3, 4),
		(6, '2019-08-05', 2, 2, 4)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Orders table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Items (item_id, item_brand) VALUES 
		(1, 'Samsung'),
		(2, 'Lenovo'),
		(3, 'LG'),
		(4, 'HP')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Items table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]struct {
		joinDate     string
		ordersIn2019 int
	}{
		1: {"2018-01-01", 1},
		2: {"2018-02-09", 2},
		3: {"2018-01-19", 0},
		4: {"2018-05-21", 0},
	}

	for rows.Next() {
		var buyerId int
		var joinDate string
		var ordersIn2019 int
		err := rows.Scan(&buyerId, &joinDate, &ordersIn2019)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedData, ok := expected[buyerId]
		if !ok || joinDate != expectedData.joinDate || ordersIn2019 != expectedData.ordersIn2019 {
			t.Errorf("Unexpected result for buyer_id: %d. got join_date: %s, orders_in_2019: %d, want join_date: %s, orders_in_2019: %d",
				buyerId, joinDate, ordersIn2019, expectedData.joinDate, expectedData.ordersIn2019)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

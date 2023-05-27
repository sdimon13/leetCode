package sales_person

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"reflect"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSalesPersonWithNoOrdersFromRed(t *testing.T) {
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

	// Drop the SalesPerson, Company and Orders table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS SalesPerson, Company, Orders`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the SalesPerson, Company and Orders table.
	_, err = db.Exec(`CREATE TABLE SalesPerson (
		sales_id int PRIMARY KEY,
		name text,
		salary int,
		commission_rate int,
		hire_date date
	)`)

	if err != nil {
		t.Fatalf("Failed to create SalesPerson table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Company (
		com_id int PRIMARY KEY,
		name text,
		city text
	)`)

	if err != nil {
		t.Fatalf("Failed to create Company table: %s", err)
	}

	_, err = db.Exec(`CREATE TABLE Orders (
		order_id int PRIMARY KEY,
		order_date date,
		com_id int,
		sales_id int,
		amount int
	)`)

	if err != nil {
		t.Fatalf("Failed to create Orders table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES 
		(1, 'John', 100000, 6, '2006-04-01'), 
		(2, 'Amy', 12000, 5, '2010-05-01'),
		(3, 'Mark', 65000, 12, '2008-12-25'),
		(4, 'Pam', 25000, 25, '2005-01-01'),
		(5, 'Alex', 5000, 10, '2007-02-03')`)

	if err != nil {
		t.Fatalf("Failed to insert data into SalesPerson table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Company (com_id, name, city) VALUES 
		(1, 'RED', 'Boston'),
		(2, 'ORANGE', 'New York'),
		(3, 'YELLOW', 'Boston'),
		(4, 'GREEN', 'Austin')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Company table: %s", err)
	}

	_, err = db.Exec(`INSERT INTO Orders (order_id, order_date, com_id, sales_id, amount) VALUES 
		(1, '2014-01-01', 3, 4, 10000),
		(2, '2014-02-01', 4, 5, 5000),
		(3, '2014-03-01', 1, 1, 50000),
		(4, '2014-04-01', 1, 4, 25000)`)

	if err != nil {
		t.Fatalf("Failed to insert data into Orders table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []string{"Amy", "Mark", "Alex"}

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		names = append(names, name)
	}

	sort.Strings(names)
	sort.Strings(expected)

	if !reflect.DeepEqual(names, expected) {
		t.Errorf("Unexpected result: got %v, want %v", names, expected)
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

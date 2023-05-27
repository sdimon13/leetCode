package calculate_special_bonus

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCalculateSpecialBonus(t *testing.T) {
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

	// Create the Employees table.
	_, err = db.Exec(`CREATE TABLE Employees (
  		employee_id int,
  		name varchar(255),
  		salary int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Employees (employee_id, name, salary) VALUES 
		(2, 'Meir', 3000), 
		(3, 'Michael', 3800), 
		(7, 'Addilyn', 7400), 
		(8, 'Juan', 6100), 
		(9, 'Kannon', 7700)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file and Verify the result.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	expected := []struct {
		employee_id int
		bonus       int
	}{
		{2, 0},
		{3, 0},
		{7, 7400},
		{8, 0},
		{9, 7700},
	}

	i := 0
	for rows.Next() {
		var employee_id, bonus int
		err := rows.Scan(&employee_id, &bonus)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if employee_id != expected[i].employee_id || bonus != expected[i].bonus {
			t.Errorf("Expected row %d to be %v, got %d %d", i+1, expected[i], employee_id, bonus)
		}
		i++
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Employees`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

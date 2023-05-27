package swap_salary

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSwapSalary(t *testing.T) {
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

	// Create the Salary table.
	_, err = db.Exec(`CREATE TABLE Salary (
  		id int,
  		name varchar(255),
  		sex ENUM('m', 'f'),
  		salary int
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Salary (id, name, sex, salary) VALUES 
		(1, 'A', 'm', 2500), 
		(2, 'B', 'f', 1500), 
		(3, 'C', 'm', 5500), 
		(4, 'D', 'f', 500)`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	_, err = db.Exec(string(query))
	if err != nil {
		t.Errorf("SQL query returned error: %s", err)
	}

	// Verify the result.
	rows, err := db.Query(`SELECT id, name, sex, salary FROM Salary ORDER BY id`)
	if err != nil {
		t.Errorf("Failed to select data: %s", err)
	}

	expected := []struct {
		id     int
		name   string
		sex    string
		salary int
	}{
		{1, "A", "f", 2500},
		{2, "B", "m", 1500},
		{3, "C", "f", 5500},
		{4, "D", "m", 500},
	}

	i := 0
	for rows.Next() {
		var id int
		var name, sex string
		var salary int
		err := rows.Scan(&id, &name, &sex, &salary)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if id != expected[i].id || name != expected[i].name || sex != expected[i].sex || salary != expected[i].salary {
			t.Errorf("Expected row %d to be %v, got %v %s %s %d", i+1, expected[i], id, name, sex, salary)
		}
		i++
	}

	// Drop the table at the end of the test.
	_, err = db.Exec(`DROP TABLE Salary`)
	if err != nil {
		t.Errorf("Failed to drop table: %s", err)
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

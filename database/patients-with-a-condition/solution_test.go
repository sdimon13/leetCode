package patients_with_a_condition

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPatientsWithCondition(t *testing.T) {
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

	// Drop the Patients table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Patients`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Patients table.
	_, err = db.Exec(`CREATE TABLE Patients (
  		patient_id int,
  		patient_name varchar(255),
		conditions varchar(255)
	)`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Patients (patient_id, patient_name, conditions) VALUES 
		(1, 'Daniel', 'YFEV COUGH'), 
		(2, 'Alice', ''), 
		(3, 'Bob', 'DIAB100 MYOP'),
		(4, 'George', 'ACNE DIAB100'), 
		(5, 'Alain', 'DIAB201')`)
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := []struct {
		patient_id   int
		patient_name string
		conditions   string
	}{
		{3, "Bob", "DIAB100 MYOP"},
		{4, "George", "ACNE DIAB100"},
	}

	i := 0
	for rows.Next() {
		var patient_id int
		var patient_name string
		var conditions string
		err := rows.Scan(&patient_id, &patient_name, &conditions)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
		}

		if patient_id != expected[i].patient_id || patient_name != expected[i].patient_name || conditions != expected[i].conditions {
			t.Errorf("Expected row %d to be %v, got %d %s %s", i+1, expected[i], patient_id, patient_name, conditions)
		}
		i++
	}

	// Close the database connection.
	err = database.CloseDB()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

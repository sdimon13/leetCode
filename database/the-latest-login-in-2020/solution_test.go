package the_latest_login_in_2020

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestLatestLoginIn2020(t *testing.T) {
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

	// Drop the Logins table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Logins`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Logins table.
	_, err = db.Exec(`CREATE TABLE Logins (
		user_id int,
		time_stamp datetime,
		PRIMARY KEY(user_id, time_stamp)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Logins table: %s", err)
	}

	// Populate the tables with initial data.
	_, err = db.Exec(`INSERT INTO Logins (user_id, time_stamp) VALUES 
		(6, '2020-06-30 15:06:07'),
		(6, '2021-04-21 14:06:06'),
		(6, '2019-03-07 00:18:15'),
		(8, '2020-02-01 05:10:53'),
		(8, '2020-12-30 00:46:50'),
		(2, '2020-01-16 02:49:50'),
		(2, '2019-08-25 07:59:08'),
		(14, '2019-07-14 09:00:00'),
		(14, '2021-01-06 11:59:59')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Logins table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]string{
		6: "2020-06-30 15:06:07",
		8: "2020-12-30 00:46:50",
		2: "2020-01-16 02:49:50",
	}

	for rows.Next() {
		var user_id int
		var last_stamp string
		err := rows.Scan(&user_id, &last_stamp)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		expectedStamp, ok := expected[user_id]
		if !ok || last_stamp != expectedStamp {
			t.Errorf("Unexpected result for user_id: %d. got last_stamp: %s, want: %s", user_id, last_stamp, expectedStamp)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

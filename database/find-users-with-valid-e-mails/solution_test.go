package find_users_with_valid_e_mails

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFindUsersWithValidEmails(t *testing.T) {
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

	// Drop the Users table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS Users`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the Users table.
	_, err = db.Exec(`CREATE TABLE Users (
		user_id int PRIMARY KEY,
		name varchar(255),
		mail varchar(255)
	)`)

	if err != nil {
		t.Fatalf("Failed to create Users table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO Users (user_id, name, mail) VALUES 
		(1, 'Winston', 'winston@leetcode.com'), 
		(2, 'Jonathan', 'jonathanisgreat'),
		(3, 'Annabelle', 'bella-@leetcode.com'),
		(4, 'Sally', 'sally.come@leetcode.com'),
		(5, 'Marwan', 'quarz#2020@leetcode.com'),
		(6, 'David', 'david69@gmail.com'),
		(7, 'Shapiro', '.shapo@leetcode.com')`)

	if err != nil {
		t.Fatalf("Failed to insert data into Users table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[int]struct{}{
		1: {},
		3: {},
		4: {},
	}

	for rows.Next() {
		var user_id int
		var name string
		var mail string
		err := rows.Scan(&user_id, &name, &mail)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		_, ok := expected[user_id]
		if !ok {
			t.Errorf("Unexpected result for user_id: %d", user_id)
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

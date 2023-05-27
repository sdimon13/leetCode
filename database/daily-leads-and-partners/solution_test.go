package daily_leads_and_partners

import (
	"github.com/sdimon13/leetCode/database"
	"io/ioutil"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDistinctLeadsAndPartnersPerDateAndMake(t *testing.T) {
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

	// Drop the DailySales table if it exists.
	_, err = db.Exec(`DROP TABLE IF EXISTS DailySales`)
	if err != nil {
		t.Fatalf("Failed to drop table: %s", err)
	}

	// Create the DailySales table.
	_, err = db.Exec(`CREATE TABLE DailySales (
		date_id date,
		make_name varchar(255),
		lead_id int,
		partner_id int
	)`)

	if err != nil {
		t.Fatalf("Failed to create DailySales table: %s", err)
	}

	// Populate the table with initial data.
	_, err = db.Exec(`INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES 
		('2020-12-08', 'toyota', 0, 1), 
		('2020-12-08', 'toyota', 1, 0),
		('2020-12-08', 'toyota', 1, 2),
		('2020-12-07', 'toyota', 0, 2),
		('2020-12-07', 'toyota', 0, 1),
		('2020-12-08', 'honda', 1, 2),
		('2020-12-08', 'honda', 2, 1),
		('2020-12-07', 'honda', 0, 1),
		('2020-12-07', 'honda', 1, 2),
		('2020-12-07', 'honda', 2, 1)
	`)

	if err != nil {
		t.Fatalf("Failed to insert data into DailySales table: %s", err)
	}

	// Execute the SQL query from the file.
	rows, err := db.Query(string(query))
	if err != nil {
		t.Fatalf("Failed to execute SQL query from solution.sql: %s", err)
	}

	expected := map[string]map[string]map[string]int{
		"2020-12-08": {
			"toyota": {"unique_leads": 2, "unique_partners": 3},
			"honda":  {"unique_leads": 2, "unique_partners": 2},
		},
		"2020-12-07": {
			"toyota": {"unique_leads": 1, "unique_partners": 2},
			"honda":  {"unique_leads": 3, "unique_partners": 2},
		},
	}

	for rows.Next() {
		var date_id string
		var make_name string
		var unique_leads int
		var unique_partners int
		err := rows.Scan(&date_id, &make_name, &unique_leads, &unique_partners)
		if err != nil {
			t.Errorf("Failed to scan row: %s", err)
			continue
		}

		if _, ok := expected[date_id]; !ok {
			t.Errorf("Unexpected date_id: %s", date_id)
			continue
		}

		if _, ok := expected[date_id][make_name]; !ok {
			t.Errorf("Unexpected make_name for date_id %s: %s", date_id, make_name)
			continue
		}

		if unique_leads != expected[date_id][make_name]["unique_leads"] {
			t.Errorf("Unexpected unique_leads for date_id %s and make_name %s. got: %d, want: %d", date_id, make_name, unique_leads, expected[date_id][make_name]["unique_leads"])
		}

		if unique_partners != expected[date_id][make_name]["unique_partners"] {
			t.Errorf("Unexpected unique_partners for date_id %s and make_name %s. got: %d, want: %d", date_id, make_name, unique_partners, expected[date_id][make_name]["unique_partners"])
		}
	}

	// Close the database connection.
	err = db.Close()
	if err != nil {
		t.Errorf("Failed to close database connection: %s", err)
	}
}

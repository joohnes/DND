package internal

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	err = createDB(db)
	return db, err
}

func createDB(db *sql.DB) error {
	var query string
	tx, err := db.Begin()
	defer tx.Rollback()

	// table players
	query = `CREATE TABLE IF NOT EXISTS players (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	level INTEGER,
	health INTEGER,
	mana INTEGER,
	strength INTEGER,
	endurance INTEGER,
	perception INTEGER,
	intelligence INTEGER,
	agility INTEGER,
	accuracy INTEGER,
	charisma INTEGER,
	class TEXT,
	race TEXT,
	subrace TEXT,
	session INTEGER
);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table items
	query = `CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		ability TEXT,
		rarity INTEGER,
		strength INTEGER,
		endurance INTEGER,
		perception INTEGER,
		intelligence INTEGER,
		agility INTEGER,
		accuracy INTEGER,
		charisma INTEGER,
		isbag BOOLEAN
	);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table player_items
	query = `CREATE TABLE IF NOT EXISTS player_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player INTEGER,
		item INTEGER
	);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table bag of holding
	query = `CREATE TABLE IF NOT EXISTS bag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item INTEGER
		);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// create basic bag of holding
	if tx.QueryRow(`SELECT * FROM items WHERE name = "Bag of Holding"`).Scan(nil) == sql.ErrNoRows {
		query = `INSERT INTO items (name, isbag) VALUES (?, ?)`
		_, _ = tx.Exec(query, "Bag of Holding", true)
	}

	// create basic player
	if tx.QueryRow(`SELECT * FROM players WHERE name = "Basic Player"`).Scan(nil) == sql.ErrNoRows {
		query = `INSERT INTO players (name, level) VALUES (?, ?)`
		_, _ = tx.Exec(query, "Basic Player", 99)
	}

	return tx.Commit()
}

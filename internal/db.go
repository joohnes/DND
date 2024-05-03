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
	_, err := db.Exec(query)
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
		charisma INTEGER
	);`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	// table player_items
	query = `CREATE TABLE IF NOT EXISTS player_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player INTEGER,
		item INTEGER
	);`
	_, err = db.Exec(query)
	return err
}

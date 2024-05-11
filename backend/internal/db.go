package internal

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	database = "database.db"

	BasicPlayer = "BasicPlayer"
)

func ConnectDB() (*sql.DB, error) {
	cfg := NewCfg()
	var db *sql.DB
	var err error

	dbCfg := mysql.NewConfig()

	if cfg.Production {
		dbCfg.User = cfg.DBUsername
		dbCfg.Passwd = cfg.DBPassword
		dbCfg.Addr = cfg.DBAdress
		dbCfg.DBName = cfg.DBName

		db, err = sql.Open("mysql", dbCfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}

		pingErr := db.Ping()
		if pingErr != nil {
			panic(pingErr)
		}
	} else {
		db, err = sql.Open("sqlite3", database)
		if err != nil {
			return nil, err
		}
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
		rarity TEXT,
		strength INTEGER,
		endurance INTEGER,
		perception INTEGER,
		intelligence INTEGER,
		agility INTEGER,
		accuracy INTEGER,
		charisma INTEGER,
		quantity INTEGER,
		equipped INTEGER
 	);`
	// TODO: equipped
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

	// table bag_items
	query = `CREATE TABLE IF NOT EXISTS bag_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item INTEGER
		);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS settings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		value TEXT
		)`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// create holder in settings
	if tx.QueryRow(`SELECT name FROM settings WHERE name = 'bag_holder';`).Scan(nil) == sql.ErrNoRows {
		_, _ = tx.Exec(`INSERT INTO settings (name, value) VALUES (?, ?)`, "bag_holder", "0")
	}

	// create basic player
	var name string
	if tx.QueryRow(`SELECT name FROM players WHERE id = 0;`).Scan(&name) == sql.ErrNoRows {
		if name != BasicPlayer {
			query = `UPDATE players SET id = id + 1;`
			_, _ = tx.Exec(query)
			query = `INSERT INTO players (id, name) VALUES (?, ?);`
			_, _ = tx.Exec(query, 0, "Basic Player")
		}
	}

	return tx.Commit()
}

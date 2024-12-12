package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	if cfg.Production {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBAdress, cfg.DBName)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		if err := db.Ping(); err != nil {
			panic(err)
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
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	createTablesMysql(tx)

	// create holder in settings
	if tx.QueryRow(`SELECT name FROM settings WHERE name = 'bag_holder';`).Scan(nil) == sql.ErrNoRows {
		_, _ = tx.Exec(`INSERT INTO settings (name, value) VALUES (?, ?)`, "bag_holder", "0")
	}

	// create basic player
	var name string
	if tx.QueryRow(`SELECT name FROM players WHERE id = 0;`).Scan(&name) != sql.ErrNoRows {
		if name != BasicPlayer {
			query := `UPDATE players SET id = id + 1;`
			_, _ = tx.Exec(query)
			query = `INSERT INTO players (id, name) VALUES (?, ?);`
			_, _ = tx.Exec(query, 0, "Basic Player")
		}
	}

	return tx.Commit()
}

func createTablesMysql(tx *sql.Tx) error {
	// table players
	query := `CREATE TABLE IF NOT EXISTS players (
		id INTEGER AUTO_INCREMENT,
		name TEXT,
		level INTEGER,
		health INTEGER DEFAULT 100,
		mana INTEGER DEFAULT 100,
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
		alcohol_level INTEGER DEFAULT 0,
		zgon BOOLEAN DEFAULT FALSE,
		PRIMARY KEY (id)
);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	// table items
	query = `CREATE TABLE IF NOT EXISTS items (
			id INTEGER AUTO_INCREMENT,
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
			attack INTEGER,
			defense INTEGER,
			permille INTEGER,
			slot INTEGER DEFAULT 0,
			PRIMARY KEY (id)
 	);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table player_items
	query = `CREATE TABLE IF NOT EXISTS player_items (
			id INTEGER AUTO_INCREMENT,
			player INTEGER,
			item INTEGER,
			equipped BOOLEAN DEFAULT FALSE,
			PRIMARY KEY (id)
		);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table bag_items
	query = `CREATE TABLE IF NOT EXISTS bag_items (
			id INTEGER AUTO_INCREMENT,
			item INTEGER,
			PRIMARY KEY (id)
			);`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	// table settings
	query = `CREATE TABLE IF NOT EXISTS settings (
			id INTEGER AUTO_INCREMENT,
			name TEXT,
			value TEXT,
			PRIMARY KEY (id)
			)`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

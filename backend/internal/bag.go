package internal

import (
	"database/sql"

	"github.com/pkg/errors"
)

// bag of holding
type Bag struct {
	Id     int    `json:"-"`
	Holder string `json:"holder"`
	Items  []Item `json:"items"`
}

func (srv *Service) GetBag() (*Bag, error) {
	var items []Item
	bag := &Bag{}

	query := "SELECT name FROM players where id=(SELECT value from settings WHERE name='bag_holder')"
	err := srv.db.QueryRow(query).Scan(&bag.Holder)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	query = "SELECT * FROM items WHERE id in (SELECT item FROM bag_items)"
	rows, err := srv.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Description,
			&i.Ability,
			&i.Rarity,
			&i.Strength,
			&i.Endurance,
			&i.Perception,
			&i.Intelligence,
			&i.Agility,
			&i.Accuracy,
			&i.Charisma,
			&i.Quantity,
			&i.Attack,
			&i.Defense,
			&i.Permille,
			&i.Slot,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	bag.Items = items

	return bag, nil
}

func (srv *Service) GetBagHolderName() (string, error) {
	query := "SELECT name FROM players where id=(SELECT value from settings WHERE name='bag_holder')"
	row := srv.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return "", err
	}
	name := ""
	err := row.Scan(&name)
	return name, err
}

func (srv *Service) ChangeBagHolder(playerID int) error {
	query := "UPDATE settings SET value=? WHERE name='bag_holder'"
	_, err := srv.db.Exec(query, playerID)
	return err
}

func (srv *Service) DropItemFromBag(itemID int) error {
	query := "DELETE FROM bag_items WHERE item=?"
	_, err := srv.db.Exec(query, itemID)
	return err
}

func (srv *Service) AddItemToBag(itemID int) error {
	query := "INSERT INTO bag_items (item) VALUES (?)"
	_, err := srv.db.Exec(query, itemID)
	return err
}

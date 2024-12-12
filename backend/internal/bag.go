package internal

import (
	"database/sql"
	"strconv"

	"github.com/pkg/errors"
)

// bag of holding
type Bag struct {
	Id     int   `json:"-"`
	Holder int   `json:"-"`
	Items  []int `json:"items"`
}

func (srv *Service) GetBagID() (int, error) {
	var id int
	query := `SELECT id FROM items where isbag = true`
	err := srv.db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (srv *Service) GetBag() (*Bag, error) {
	var items []int
	bag := &Bag{}
	var tmp string
	var holder int

	query := "SELECT value FROM settings WHERE name='bag_holder'"
	err := srv.db.QueryRow(query).Scan(&tmp)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	holder, err = strconv.Atoi(tmp)
	if err != nil {
		return nil, err
	}

	query = "SELECT item FROM bag_items"
	rows, err := srv.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int
		if err := rows.Scan(&itemID); err != nil {
			return nil, err
		}
		items = append(items, itemID)
	}
	bag.Holder = holder
	bag.Items = items

	return bag, nil
}

func (bag *Bag) IsItemInBag(id int) bool {
	for _, item := range bag.Items {
		if item == id {
			return true
		}
	}
	return false
}

func (srv *Service) GetBagHolderName(id int) (string, error) {
	query := "SELECT name FROM players where id=?"
	row := srv.db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return "", err
	}
	var name string
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

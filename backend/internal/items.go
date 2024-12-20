package internal

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Ability      string `json:"ability"`
	Rarity       string `json:"rarity"`
	Strength     int    `json:"strength"`
	Endurance    int    `json:"endurance"`
	Perception   int    `json:"perception"`
	Intelligence int    `json:"intelligence"`
	Agility      int    `json:"agility"`
	Accuracy     int    `json:"accuracy"`
	Charisma     int    `json:"charisma"`
	Quantity     int    `json:"quantity"`
	Attack       int    `json:"attack"`
	Defense      int    `json:"defense"`
	Permille     int    `json:"permille"` // promile
	Slot         int    `json:"slot"`
	Owner        string `json:"owner"`
}

func IsValidRarity(rarity string) error {
	switch rarity {
	case "Common", "Rare", "Epic", "Legendary", "Artefact":
		return nil
	}
	return ErrInvalidRarity
}

func (srv *Service) getOwner(itemID int) (string, error) {
	query := "SELECT name FROM players where id = (SELECT player FROM player_items WHERE item = ?"
	row := srv.db.QueryRow(query, itemID)
	if err := row.Err(); err != nil {
		return "", errors.Wrap(err, "query row")
	}
	var name string
	if err := row.Scan(&name); err != nil {
		return "", errors.Wrap(err, "scan")
	}

	return name, nil
}

func (srv *Service) GetItemByID(id int) (*Item, error) {
	query := "SELECT * FROM items WHERE id=?"
	row := srv.db.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return nil, errors.Wrap(err, "queryrow")
	}

	i := &Item{}
	err := row.Scan(
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
	)
	if err != nil {
		return nil, errors.Wrap(err, "scan")
	}

	return i, nil
}

func (srv *Service) GetPlayerItems(playerID int) ([]*Item, error) {
	query := "SELECT * FROM items WHERE id in (SELECT item FROM player_items WHERE player=?)"
	rows, err := srv.db.Query(query, playerID)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	items := make([]*Item, 0)
	for rows.Next() {
		i := &Item{}
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, errors.Wrap(err, "scan")
		}

		items = append(items, i)
	}

	return items, nil
}

func (srv *Service) GetItems() ([]*Item, error) {
	query := "SELECT * FROM items WHERE id not in (SELECT item FROM bag_items)"
	rows, err := srv.db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "query")
	}
	defer rows.Close()

	var items []*Item
	for rows.Next() {
		i := &Item{}
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, errors.Wrap(err, "scan")
		}

		items = append(items, i)
	}

	owners := make(map[int]string, 0)
	query = "SELECT i.item, p.name FROM player_items i INNER JOIN players p on i.player = p.id"
	rowsOwners, err := srv.db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "query")
	}
	defer rowsOwners.Close()

	for rowsOwners.Next() {
		var id int
		var name string
		if err := rowsOwners.Scan(&id, &name); err != nil {
			return nil, errors.Wrap(err, "scan owners")
		}

		owners[id] = name
	}

	for _, item := range items {
		if owner, exist := owners[item.Id]; exist {
			item.Owner = owner
		}
	}

	return items, nil
}

func (srv *Service) CreateItem(i Item) (*Item, error) {
	query := "INSERT INTO items (name, description, ability, rarity, strength, endurance, perception, intelligence, agility, accuracy, charisma, quantity, attack, defense, permille, slot) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := srv.db.Exec(query, i.Name, i.Description, i.Ability, i.Rarity, i.Strength, i.Endurance, i.Perception, i.Intelligence, i.Agility, i.Accuracy, i.Charisma, i.Quantity, i.Attack, i.Defense, i.Permille, i.Slot)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	i.Id = int(id)
	return &i, err
}

func (srv *Service) UpdateItem(i Item) error {
	query := `UPDATE items
				SET
					name=?,
					description=?,
					ability=?,
					rarity=?,
					strength=?,
					endurance=?,
					perception=?,
					intelligence=?,
					agility=?,
					accuracy=?,
					charisma=?,
					quantity=?,
					attack=?,
					defense=?,
					permille=?,
					slot=?
				WHERE
					id=?;
	`
	_, err := srv.db.Exec(query, i.Name, i.Description, i.Ability, i.Rarity, i.Strength, i.Endurance, i.Perception, i.Intelligence, i.Agility, i.Accuracy, i.Charisma, i.Quantity, i.Attack, i.Defense, i.Permille, i.Slot, i.Id)
	return err
}

func (srv *Service) DeleteItem(id int) error {
	_, err := srv.db.Exec("DELETE FROM items WHERE id=?", id)
	if err != nil {
		return err
	}

	_, err = srv.db.Exec("DELETE FROM player_items WHERE item=?", id)
	return err
}

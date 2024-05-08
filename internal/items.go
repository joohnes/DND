package internal

import "github.com/pkg/errors"

type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Ability      string `json:"ability"`
	Rarity       int    `json:"rarity"`
	Strength     int    `json:"strength"`
	Endurance    int    `json:"endurance"`
	Perception   int    `json:"perception"`
	Intelligence int    `json:"intelligence"`
	Agility      int    `json:"agility"`
	Accuracy     int    `json:"accuracy"`
	Charisma     int    `json:"charisma"`
	Quantity     int    `json:"quantity"`
	Equipped     int    `json:"equipped"`
}

func (srv *Service) GetItemByID(id int) *Item {
	for _, item := range srv.items {
		if item.Id == id {
			return item
		}
	}
	return nil
}

func (srv *Service) GetItemsIDs() []int {
	var ids []int
	for _, i := range srv.items {
		if srv.IsInBag(i.Id) {
			continue
		}
		ids = append(ids, i.Id)
	}
	return ids
}

func (srv *Service) CreateItem(i Item) (*Item, error) {
	query := "INSERT INTO items (name, description, ability, rarity, strength, endurance, perception, intelligence, agility, accuracy, charisma, quantity, equipped, isbag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := srv.db.Exec(query, i.Name, i.Description, i.Ability, i.Rarity, i.Strength, i.Endurance, i.Perception, i.Intelligence, i.Agility, i.Accuracy, i.Charisma, i.Quantity, false, false)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	i.Id = int(id)
	srv.items = append(srv.items, &i)
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
					quantity=?
				WHERE
					id=?;
	`
	_, err := srv.db.Exec(query, i.Name, i.Description, i.Ability, i.Rarity, i.Strength, i.Endurance, i.Perception, i.Intelligence, i.Agility, i.Accuracy, i.Charisma, i.Quantity, i.Id)
	if err != nil {
		return errors.Wrap(err, "failed to update item")
	}
	return srv.ResetObjects(ItemType)
}

func (srv *Service) EquipItem(itemID int, slot int) error {
	query := "UPDATE items SET equipped = ? WHERE id = ?"
	_, err := srv.db.Exec(query, slot, itemID)
	if err != nil {
		return err
	}

	return srv.ResetObjects(ItemType)
}

func (srv *Service) UnequipItem(itemID int) error {
	query := "UPDATE items SET equipped = 0 WHERE id = ?"
	_, err := srv.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	return srv.ResetObjects(ItemType)
}

func (srv *Service) GetItemsFromDB() ([]*Item, error) {
	query := "SELECT id, name, description, ability, rarity, strength, endurance, perception, intelligence, agility, accuracy, charisma, quantity FROM items"
	rows, err := srv.db.Query(query)
	defer func() { rows.Close() }()
	if err != nil {
		return nil, err
	}

	playerItems := make(map[int]int)
	query = "SELECT player, item FROM player_items"
	rowsPlayerItems, err := srv.db.Query(query)
	defer func() { rowsPlayerItems.Close() }()
	if err != nil {
		return nil, err
	}
	for rowsPlayerItems.Next() {
		var player, item int
		rowsPlayerItems.Scan(&player, &item)
		playerItems[item] = player
	}

	var items []*Item
	for rows.Next() {
		i := &Item{}
		rows.Scan(
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
		)

		i.Owner = playerItems[i.Id]
		items = append(items, i)
	}

	return items, nil
}

func (srv *Service) DeleteItem(id int) error {
	_, err := srv.db.Exec("DELETE FROM items WHERE id=?", id)
	if err != nil {
		return err
	}

	_, err = srv.db.Exec("DELETE FROM player_items WHERE item=?", id)
	if err != nil {
		return err
	}
	return srv.ResetObjects(ItemType)
}

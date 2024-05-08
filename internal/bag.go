package internal

// bag of holding
type Bag struct {
	Id    int   `json:"-"`
	Owner int   `json:"owner"`
	Items []int `json:"items"`
}

func (srv *Service) GetBagOwner() int {
	for _, player := range srv.players {
		for _, item := range player.Items {
			if item == srv.bag.Id {
				return player.Id
			}
		}
	}
	return 0
}

func (srv *Service) GetBagFromDB() (*Bag, error) {
	query := "SELECT item FROM bag_items"
	rows, err := srv.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int
	bag := &Bag{}

	for rows.Next() {
		var itemID int
		if err := rows.Scan(&itemID); err != nil {
			return nil, err
		}
		items = append(items, itemID)
	}
	bag.Items = items
	// get bag owner id
	bag.Owner = srv.GetBagOwner()
	return bag, nil
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

func (srv *Service) GetBag() *Bag {
	return srv.bag
}

func (bag *Bag) IsItemInBag(id int) bool {
	for _, item := range bag.Items {
		if item == id {
			return true
		}
	}
	return false
}

func (srv *Service) DropItemFromBag(itemID int) error {
	if !srv.bag.IsItemInBag(itemID) {
		return ErrItemNotFound
	}
	query := "DELETE FROM bag_items WHERE item=?"
	_, err := srv.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	// bag := srv.GetBag()
	// for i := 0; i < len(bag.Items); i++ {
	// 	if bag.Items[i] == itemID {
	// 		bag.Items[i] = bag.Items[len(bag.Items)-1]
	// 		bag.Items = bag.Items[:len(bag.Items)-1]
	// 	}
	// }
	return srv.ResetObjects(BagType)
}

func (srv *Service) AddItemToBag(itemID int) error {
	if srv.bag.IsItemInBag(itemID) {
		return ErrItemAlreadyExists
	}
	query := "INSERT INTO bag_items (item) VALUES (?)"
	_, err := srv.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	// bag := srv.GetBag()
	// bag.Items = append(bag.Items, id)

	return srv.ResetObjects(BagType)
}

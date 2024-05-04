package internal

// bag of holding
type Bag struct {
	Owner int     `json:"owner"`
	Items []*Item `json:"items"`
}

func (srv *Service) GetBagOwner() int {
	for _, player := range srv.players {
		for _, item := range player.Items {
			if item.Name == "Bag of Holding" {
				return player.Id
			}
		}
	}
	return 0
}

func (srv *Service) GetBagFromDB() (*Bag, error) {
	query := "SELECT item FROM bag"
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

	for _, itemID := range items {
		i := srv.GetItemByID(itemID)
		bag.Items = append(bag.Items, i)
	}

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

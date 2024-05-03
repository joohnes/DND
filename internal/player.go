package internal

import (
	"database/sql"
	"errors"
)

type Stat string
type Class string

type Player struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Level        int     `json:"level"`
	Health       int     `json:"health"`
	Mana         int     `json:"mana"`
	Class        string  `json:"class"`
	Race         string  `json:"race"`
	Subrace      string  `json:"subrace"`
	Strength     int     `json:"strength"`
	Endurance    int     `json:"endurance"`
	Perception   int     `json:"perception"`
	Intelligence int     `json:"intelligence"`
	Agility      int     `json:"agility"`
	Accuracy     int     `json:"accuracy"`
	Charisma     int     `json:"charisma"`
	Session      int     `json:"session"`
	Items        []*Item `json:"items"`
}

func (p *Player) GetItems() []*Item {
	return p.Items
}

func (srv *Service) DropItem(playerID, itemID int) error {
	p := srv.GetPlayerByID(playerID)
	for i := 0; i < len(p.Items); i++ {
		if p.Items[i].Id == itemID {
			p.Items[i] = p.Items[len(p.Items)-1]
			p.Items = p.Items[:len(p.Items)-1]
		}
	}
	srv.GetItemByID(itemID).Owner = 0

	query := "DELETE FROM player_items WHERE player=? AND item=?"
	_, err := srv.db.Exec(query, playerID, itemID)
	return err
}

func (srv *Service) AddItem(playerID, itemID int) error {
	_, err := srv.db.Exec("INSERT INTO player_items (player, item) VALUES (?, ?)", playerID, itemID)
	if err != nil {
		return err
	}
	item := srv.GetItemByID(itemID)
	player := srv.GetPlayerByID(playerID)
	if player == nil {
		return errors.New("player with given id not found")
	}
	player.Items = append(player.Items, item)
	srv.GetItemByID(itemID).Owner = playerID

	return nil
}

func (srv *Service) GetPlayers() []*Player {
	return srv.players
}

func (srv *Service) GetPlayerByID(id int) *Player {
	for _, player := range srv.players {
		if player.Id == id {
			return player
		}
	}
	return nil
}

func (srv *Service) CreatePlayer(p Player) (*Player, error) {
	query := "INSERT INTO players (name, level, class, race, subrace, strength, endurance, perception, intelligence, agility, accuracy, charisma, health, mana, session) VALUES (?, ?, ?,?, ?, ?,?, ?, ?,?, ?, ?, 100, 100, ?)"
	res, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma, p.Session)
	id, _ := res.LastInsertId()
	p.Id = int(id)
	return &p, err
}

func (srv *Service) UpdatePlayer(p Player) error {
	query := `UPDATE players
				SET
					name=?,
					level=?,
					health=?,
					mana=?,
					class=?,
					race=?,
					subrace=?,
					strength=?,
					endurance=?,
					perception=?,
					intelligence=?,
					agility=?,
					accuracy=?,
					charisma=?,
					session=?
				WHERE
					id=?;
	`
	_, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma, p.Session, p.Id)
	return err
}

func (srv *Service) DeletePlayer(id int) error {
	_, err := srv.db.Exec("DELETE FROM players WHERE id=?", id)
	if err != nil {
		return err
	}
	_, err = srv.db.Exec("DELETE FROM players WHERE player=?", id)
	return err
}

func (srv *Service) GetPlayersFromDB() ([]*Player, error) {
	query := "SELECT * FROM players WHERE name != 'Basic Player'"
	rows, err := srv.db.Query(query)
	defer func() { rows.Close() }()
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}

	players := make([]*Player, 0)
	for rows.Next() {
		p := &Player{}
		if err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Level,
			&p.Health,
			&p.Mana,
			&p.Strength,
			&p.Endurance,
			&p.Perception,
			&p.Intelligence,
			&p.Agility,
			&p.Accuracy,
			&p.Charisma,
			&p.Class,
			&p.Race,
			&p.Subrace,
			&p.Session,
		); err != nil {
			return nil, err
		}

		rows, err := srv.db.Query("SELECT * FROM items WHERE id in (SELECT item FROM player_items WHERE player = ?)", p.Id)
		defer func() { rows.Close() }()
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
			return nil, nil
		}

		for rows.Next() {
			i := &Item{}
			if err := rows.Scan(
				&i.Id,
				&i.Name,
				&i.Description,
				&i.Ability,
				&i.Strength,
				&i.Endurance,
				&i.Perception,
				&i.Intelligence,
				&i.Agility,
				&i.Accuracy,
				&i.Charisma,
			); err != nil {
				return nil, err
			}
			p.Items = append(p.Items, i)
		}
		players = append(players, p)
	}
	return players, nil
}

package internal

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Player struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Level        int    `json:"level"`
	Health       int    `json:"health"`
	Mana         int    `json:"mana"`
	Class        string `json:"class"`
	Race         string `json:"race"`
	Subrace      string `json:"subrace"`
	Strength     int    `json:"strength"`
	Endurance    int    `json:"endurance"`
	Perception   int    `json:"perception"`
	Intelligence int    `json:"intelligence"`
	Agility      int    `json:"agility"`
	Accuracy     int    `json:"accuracy"`
	Charisma     int    `json:"charisma"`
	Session      int    `json:"-"`
	Items        []int  `json:"items"`
}

type HPMana struct {
	HP   int `json:"hp"`
	Mana int `json:"mana"`
}

func (p *Player) GetItems() []int {
	return p.Items
}

func (srv *Service) DropItem(itemID int) error {
	var playerID int
	for _, p := range srv.players {
		for i := 0; i < len(p.Items); i++ {
			if p.Items[i] == itemID {
				playerID = p.Id
				p.Items[i] = p.Items[len(p.Items)-1]
				p.Items = p.Items[:len(p.Items)-1]
			}
		}
	}

	query := "DELETE FROM player_items WHERE player=? AND item=?"
	_, err := srv.db.Exec(query, playerID, itemID)
	return errors.Wrap(err, "failed to drop item")
}

func (srv *Service) AddItem(playerID, itemID int) error {
	_, err := srv.db.Exec("INSERT INTO player_items (player, item) VALUES (?, ?)", playerID, itemID)
	if err != nil {
		return err
	}

	player := srv.GetPlayerByID(playerID)
	if player == nil {
		return errors.New("player with given id not found")
	}
	player.Items = append(player.Items, itemID)

	return nil
}

func (srv *Service) GetPlayersIDs() []int {
	var ids []int
	for _, p := range srv.players {
		ids = append(ids, p.Id)
	}
	return ids
}

func (srv *Service) GetPlayerIdsWithNames() map[int]string {
	players := make(map[int]string)
	for _, player := range srv.players {
		players[player.Id] = player.Name
	}
	return players
}

func (srv *Service) GetPlayerByID(id int) *Player {
	for _, player := range srv.players {
		if player.Id == id {
			return player
		}
	}
	return nil
}

func (srv *Service) ChangeHPandMana(id int, hpmana HPMana) error {
	query := "UPDATE players SET health=?, mana=? WHERE id=?"
	_, err := srv.db.Exec(query, hpmana.HP, hpmana.Mana, id)
	if err != nil {
		return errors.Wrap(err, "query: failed to change hp and mana")
	}
	return errors.Wrap(srv.ResetObjects(PlayerType), "failed to change hp and mana")
}

func (srv *Service) CreatePlayer(p Player) (*Player, error) {
	query := "INSERT INTO players (name, level, class, race, subrace, strength, endurance, perception, intelligence, agility, accuracy, charisma, health, mana, session) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 100, 100, ?)"
	res, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma, p.Session)
	id, _ := res.LastInsertId()
	p.Id = int(id)
	p.Health = 100
	p.Mana = 100
	srv.players = append(srv.players, &p)
	return &p, err
}

func (srv *Service) UpdatePlayer(p Player) error {
	query := `UPDATE players
				SET
					name=?,
					level=?,
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
	if err != nil {
		return errors.Wrap(err, "failed to update player")
	}
	return srv.ResetObjects(PlayerType)
}

func (srv *Service) DeletePlayer(id int) error {
	_, err := srv.db.Exec("DELETE FROM players WHERE id=?", id)
	if err != nil {
		return err
	}
	_, err = srv.db.Exec("DELETE FROM player_items WHERE player=?", id)
	if err != nil {
		return err
	}
	return srv.ResetObjects(PlayerType)
}

func (srv *Service) GetPlayersFromDB() ([]*Player, error) {
	query := "SELECT * FROM players WHERE name != 'Basic Player' AND id != 0"
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

		rows, err := srv.db.Query("SELECT item FROM player_items WHERE player = ?", p.Id)
		defer func() { rows.Close() }()
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
			return nil, nil
		}

		for rows.Next() {
			var itemID int
			if err := rows.Scan(
				&itemID,
			); err != nil {
				return nil, err
			}
			p.Items = append(p.Items, itemID)
		}
		players = append(players, p)
	}
	return players, nil
}

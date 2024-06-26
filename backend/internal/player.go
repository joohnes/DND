package internal

import (
	"database/sql"
	"fmt"

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
	AlcoholLevel int    `json:"alcohol_level"`
	Zgon         bool   `json:"zgon"`
	Equipped     []int  `json:"equipped"`
	Items        []int  `json:"items"`
}

type PlayerResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Level        int    `json:"level"`
	Health       int    `json:"health"`
	Mana         int    `json:"mana"`
	Class        string `json:"class"`
	Race         string `json:"race"`
	Subrace      string `json:"subrace"`
	Strength     string `json:"strength"`
	Endurance    string `json:"endurance"`
	Perception   string `json:"perception"`
	Intelligence string `json:"intelligence"`
	Agility      string `json:"agility"`
	Accuracy     string `json:"accuracy"`
	Charisma     string `json:"charisma"`
	AlcoholLevel int    `json:"alcohol_level"`
	Zgon         bool   `json:"zgon"`
	Equipped     []int  `json:"equipped"`
	Items        []int  `json:"items"`
}

type HPMana struct {
	HP   int `json:"hp"`
	Mana int `json:"mana"`
}

type ItemsResponse struct {
	Items    []int `json:"items"`
	Equipped []int `json:"equipped"`
}

// slots - 1 head, 2 chest, 3 shoulders, 4 hands, 5 legs, 6 feet
func (p *Player) GetItems() ItemsResponse {
	return ItemsResponse{Items: p.Items, Equipped: p.Equipped}
}

func (p *Player) IsEquipped(id int) bool {
	for _, i := range p.Equipped {
		if i == id {
			return true
		}
	}
	return false
}

func (p *Player) IsInEquipment(id int) bool {
	for _, i := range p.Items {
		if i == id {
			return true
		}
	}
	return false
}

func (srv *Service) EquipItem(playerID, itemID int) error {
	fmt.Println(playerID, itemID)
	if srv.GetPlayerByID(playerID) == nil {
		return ErrNoPlayer
	} else {
		if srv.GetPlayerByID(playerID).IsEquipped(itemID) {
			return ErrAlreadyEquipped
		}
	}

	query := "UPDATE player_items SET equipped = 1 WHERE item = ?"
	_, err := srv.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	return srv.ResetObjects(PlayerType)
}

func (srv *Service) UnequipItem(itemID int) error {
	query := "UPDATE player_items SET equipped = 0 WHERE item = ?"
	_, err := srv.db.Exec(query, itemID)
	if err != nil {
		return err
	}

	return srv.ResetObjects(PlayerType)
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
		for i := 0; i < len(p.Equipped); i++ {
			if p.Equipped[i] == itemID {
				playerID = p.Id
				p.Equipped[i] = p.Equipped[len(p.Equipped)-1]
				p.Equipped = p.Equipped[:len(p.Equipped)-1]
			}
		}
	}

	query := "DELETE FROM player_items WHERE player=? AND item=?"
	_, err := srv.db.Exec(query, playerID, itemID)
	return errors.Wrap(err, "failed to drop item")
}

func (srv *Service) AddItem(playerID, itemID int) error {
	if srv.GetPlayerByID(playerID).IsInEquipment(itemID) {
		return errors.Wrap(ErrItemAlreadyExists, "eq")
	}

	_, err := srv.db.Exec("INSERT INTO player_items (player, item) VALUES (?, ?)", playerID, itemID)
	if err != nil {
		return err
	}

	player := srv.GetPlayerByID(playerID)
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

func (srv *Service) GetPlayerResponseByID(id int) *PlayerResponse {
	player := srv.GetPlayerByID(id)
	playerTotal := *player
	for _, id := range player.Items {
		item, err := srv.GetItemByID(id)
		if err != nil {
			continue
		}
		playerTotal.Strength += item.Strength
		playerTotal.Endurance += item.Endurance
		playerTotal.Perception += item.Perception
		playerTotal.Intelligence += item.Intelligence
		playerTotal.Agility += item.Agility
		playerTotal.Accuracy += item.Accuracy
		playerTotal.Charisma += item.Charisma
	}

	return &PlayerResponse{
		Id:           player.Id,
		Name:         player.Name,
		Health:       player.Health,
		Mana:         player.Mana,
		Level:        player.Level,
		Class:        player.Class,
		Race:         player.Race,
		Subrace:      player.Subrace,
		Strength:     fmt.Sprintf("%d (%d)", player.Strength, playerTotal.Strength),
		Endurance:    fmt.Sprintf("%d (%d)", player.Endurance, playerTotal.Endurance),
		Perception:   fmt.Sprintf("%d (%d)", player.Perception, playerTotal.Perception),
		Intelligence: fmt.Sprintf("%d (%d)", player.Intelligence, playerTotal.Intelligence),
		Agility:      fmt.Sprintf("%d (%d)", player.Agility, playerTotal.Agility),
		Accuracy:     fmt.Sprintf("%d (%d)", player.Accuracy, playerTotal.Accuracy),
		Charisma:     fmt.Sprintf("%d (%d)", player.Charisma, playerTotal.Charisma),
		AlcoholLevel: player.AlcoholLevel,
		Zgon:         player.Zgon,
		Equipped:     player.Equipped,
		Items:        player.Items,
	}
}

func (srv *Service) ToggleZgon(playerID int) error {
	if p := srv.GetPlayerByID(playerID); p == nil {
		return errors.Wrap(ErrNoPlayer, "toggle zgon")
	}

	query := "UPDATE players SET zgon = NOT zgon WHERE id = ?"
	_, err := srv.db.Exec(query, playerID)
	if err != nil {
		return errors.Wrap(err, "failed to toggle zgon")
	}

	return errors.Wrap(srv.ResetObjects(PlayerType), "toggle zgon")
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
	query := "INSERT INTO players (name, level, class, race, subrace, strength, endurance, perception, intelligence, agility, accuracy, charisma, health, mana) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 100, 100)"
	res, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma)
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
					alcohol_level=?
				WHERE
					id=?;
	`
	_, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma, p.AlcoholLevel, p.Id)
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
			&p.AlcoholLevel,
			&p.Zgon,
		); err != nil {
			return nil, err
		}

		rowsItems, err := srv.db.Query("SELECT item, equipped FROM player_items WHERE player = ?", p.Id)
		defer rows.Close()
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
			return nil, nil
		}

		for rowsItems.Next() {
			var itemID int
			var equipped int
			if err := rowsItems.Scan(
				&itemID,
				&equipped,
			); err != nil {
				return nil, err
			}
			if equipped == 1 {
				p.Equipped = append(p.Equipped, itemID)
				continue
			}
			p.Items = append(p.Items, itemID)
		}

		players = append(players, p)
	}
	return players, nil
}

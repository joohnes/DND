package internal

import (
	"fmt"
	"log/slog"

	"github.com/pkg/errors"
)

type Player struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Health  int    `json:"health"`
	Mana    int    `json:"mana"`
	Class   string `json:"class"`
	Race    string `json:"race"`
	Subrace string `json:"subrace"`
	Stats
	AlcoholLevel int    `json:"alcohol_level"`
	Zgon         bool   `json:"zgon"`
	Items        []Item `json:"items"`
	Equipped     []Item `json:"equipped"`
}

type PlayerResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Health  int    `json:"health"`
	Mana    int    `json:"mana"`
	Class   string `json:"class"`
	Race    string `json:"race"`
	Subrace string `json:"subrace"`
	Stats
	StatsPrintable
	AlcoholLevel int    `json:"alcohol_level"`
	Zgon         bool   `json:"zgon"`
	Equipped     []Item `json:"equipped"`
	Items        []Item `json:"items"`
	IsHolder     bool   `json:"is_holder"`
}

type Stats struct {
	Subrace      int `json:"subrace"`
	Strength     int `json:"strength"`
	Endurance    int `json:"endurance"`
	Perception   int `json:"perception"`
	Intelligence int `json:"intelligence"`
	Agility      int `json:"agility"`
	Accuracy     int `json:"accuracy"`
	Charisma     int `json:"charisma"`
}

type StatsPrintable struct {
	StrengthPrintable     string `json:"strength_printable"`
	EndurancePrintable    string `json:"endurance_printable"`
	PerceptionPrintable   string `json:"perception_printable"`
	IntelligencePrintable string `json:"intelligence_printable"`
	AgilityPrintable      string `json:"agility_printable"`
	AccuracyPrintable     string `json:"accuracy_printable"`
	CharismaPrintable     string `json:"charisma_printable"`
}

type HPMana struct {
	HP   int `json:"hp"`
	Mana int `json:"mana"`
}

type ItemsResponse struct {
	Items    []Item `json:"items"`
	Equipped []Item `json:"equipped"`
}

func (srv *Service) EquipItem(playerID, itemID int) error {
	query := "UPDATE player_items SET equipped = 1 WHERE item = ?"
	_, err := srv.db.Exec(query, itemID)
	return err
}

func (srv *Service) UnequipItem(itemID int) error {
	query := "UPDATE player_items SET equipped = 0 WHERE item = ?"
	_, err := srv.db.Exec(query, itemID)
	return err
}

func (srv *Service) DropItem(itemID int) error {
	query := "DELETE FROM player_items WHERE item=?"
	_, err := srv.db.Exec(query, itemID)
	return errors.Wrap(err, "failed to drop item")
}

func (srv *Service) AddItem(playerID, itemID int) error {
	_, err := srv.db.Exec("INSERT INTO player_items (player, item) VALUES (?, ?)", playerID, itemID)
	if err != nil {
		return err
	}

	return nil
}

func (srv *Service) GetPlayersIDs() ([]int, error) {
	query := "SELECT id FROM players"
	rows, err := srv.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "GetPlayersIDs")
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "scan")
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (srv *Service) GetPlayerIdsWithNames() (map[int]string, error) {
	query := "SELECT id, name FROM players"
	rows, err := srv.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "GetPlayersIDs")
	}
	defer rows.Close()
	idsnames := make(map[int]string, 0)

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, errors.Wrap(err, "scan")
		}
		idsnames[id] = name
	}

	return idsnames, nil
}

func (srv *Service) GetPlayerByID(id int) (*Player, error) {
	query := "SELECT * FROM players WHERE id=?"
	row := srv.db.QueryRow(query, id)
	if row.Err() != nil {
		return nil, errors.Wrap(row.Err(), "GetPlayerByID")
	}

	p := &Player{}
	err := row.Scan(
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
	)
	if err != nil {
		return nil, errors.Wrap(err, "scan")
	}

	query = "SELECT * FROM items WHERE id in (SELECT item FROM player_items WHERE player = ?)"
	rows, err := srv.db.Query(query, p.Id)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "query items")
	}

	itemsDB := make([]Item, 0)
	for rows.Next() {
		iDB := Item{}
		err := rows.Scan(
			&iDB.Id,
			&iDB.Name,
			&iDB.Description,
			&iDB.Ability,
			&iDB.Rarity,
			&iDB.Strength,
			&iDB.Endurance,
			&iDB.Perception,
			&iDB.Intelligence,
			&iDB.Agility,
			&iDB.Accuracy,
			&iDB.Charisma,
			&iDB.Quantity,
			&iDB.Attack,
			&iDB.Defense,
			&iDB.Permille,
			&iDB.Slot,
		)

		if err != nil {
			return nil, errors.Wrap(err, "scan item")
		}

		itemsDB = append(itemsDB, iDB)
	}

	query = "SELECT item FROM player_items WHERE player = ? AND equipped = 1"
	rows, err = srv.db.Query(query, p.Id)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "query equipped")
	}

	equippedIDs := make([]int, 0)
	for rows.Next() {
		id := 0
		err := rows.Scan(&id)
		if err != nil {
			return nil, errors.Wrap(err, "scan equipped")
		}

		equippedIDs = (append(equippedIDs, id))
	}

	for i, item := range itemsDB {
		if in(item.Id, equippedIDs) {
			p.Equipped = append(p.Equipped, item)
			if i >= len(itemsDB) {
				itemsDB = itemsDB[:i]
			} else {
				itemsDB = append(itemsDB[:i], itemsDB[i+1:]...)
			}
		}
	}

	p.Items = itemsDB

	return p, nil
}

func in(item int, lst []int) bool {
	for _, i := range lst {
		if i == item {
			return true
		}
	}

	return false
}

func (srv *Service) GetPlayerResponseByID(id int) (*PlayerResponse, error) {
	player, err := srv.GetPlayerByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "GetPlayerResponseByID")
	}

	playerTotal := *player
	for _, i := range player.Items {
		item, err := srv.GetItemByID(i.Id)
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

	isHolder := false
	holder, err := srv.GetBagHolderName()
	if err != nil {
		slog.Error("GetPlayerResponseByID", slog.Any("err", err))
	} else {
		isHolder = holder == player.Name
	}

	return &PlayerResponse{
		Id:      player.Id,
		Name:    player.Name,
		Health:  player.Health,
		Mana:    player.Mana,
		Level:   player.Level,
		Class:   player.Class,
		Race:    player.Race,
		Subrace: player.Subrace,
		Stats: Stats{
			Strength:     player.Strength,
			Endurance:    player.Endurance,
			Perception:   player.Perception,
			Intelligence: player.Intelligence,
			Agility:      player.Agility,
			Accuracy:     player.Accuracy,
			Charisma:     player.Charisma,
		},
		StatsPrintable: StatsPrintable{
			StrengthPrintable:     fmt.Sprintf("%d (%d)", player.Strength, playerTotal.Strength),
			EndurancePrintable:    fmt.Sprintf("%d (%d)", player.Endurance, playerTotal.Endurance),
			PerceptionPrintable:   fmt.Sprintf("%d (%d)", player.Perception, playerTotal.Perception),
			IntelligencePrintable: fmt.Sprintf("%d (%d)", player.Intelligence, playerTotal.Intelligence),
			AgilityPrintable:      fmt.Sprintf("%d (%d)", player.Agility, playerTotal.Agility),
			AccuracyPrintable:     fmt.Sprintf("%d (%d)", player.Accuracy, playerTotal.Accuracy),
			CharismaPrintable:     fmt.Sprintf("%d (%d)", player.Charisma, playerTotal.Charisma),
		},
		AlcoholLevel: player.AlcoholLevel,
		Zgon:         player.Zgon,
		Equipped:     player.Equipped,
		Items:        player.Items,
		IsHolder:     isHolder,
	}, nil
}

func (srv *Service) ToggleZgon(playerID int) error {
	query := "UPDATE players SET zgon = NOT zgon WHERE id = ?"
	_, err := srv.db.Exec(query, playerID)
	return err
}

func (srv *Service) ChangeHPandMana(id int, hpmana HPMana) error {
	query := "UPDATE players SET health=?, mana=? WHERE id=?"
	_, err := srv.db.Exec(query, hpmana.HP, hpmana.Mana, id)
	return err
}

func (srv *Service) CreatePlayer(p Player) (*Player, error) {
	query := "INSERT INTO players (name, level, class, race, subrace, strength, endurance, perception, intelligence, agility, accuracy, charisma, health, mana) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 100, 100)"
	res, err := srv.db.Exec(query, p.Name, p.Level, p.Class, p.Race, p.Subrace, p.Strength, p.Endurance, p.Perception, p.Intelligence, p.Agility, p.Accuracy, p.Charisma)
	id, _ := res.LastInsertId()
	p.Id = int(id)
	p.Health = 100
	p.Mana = 100
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
	return err
}

func (srv *Service) DeletePlayer(id int) error {
	_, err := srv.db.Exec("DELETE FROM players WHERE id=?", id)
	if err != nil {
		return err
	}
	_, err = srv.db.Exec("DELETE FROM player_items WHERE player=?", id)
	return err
}

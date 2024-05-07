package internal

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Service struct {
	db      *sql.DB
	players []*Player
	items   []*Item
	bag     *Bag
}

func NewService(db *sql.DB) (*Service, error) {
	srv := &Service{db: db}
	players, err := srv.GetPlayersFromDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetPlayersFromDB")
	}
	items, err := srv.GetItemsFromDB()
	if err != nil {
		return nil, errors.Wrap(err, "GetItemsFromDB")
	}

	srv.items = items
	srv.players = players
	srv.bag, err = srv.GetBagFromDB()
	return srv, err
}

func (srv *Service) TestObjects() {
	srv.players = append(srv.players, &Player{
		Id:           1,
		Name:         "Test",
		Level:        3,
		Health:       100,
		Mana:         100,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Class:        "Warrior",
		Race:         "zjeb",
		Subrace:      "zjeb",
		Session:      1,
		Items:        []int{69, 1337, 2137},
	})
	srv.players = append(srv.players, &Player{
		Id:           2,
		Name:         "Test2",
		Level:        69,
		Health:       69,
		Mana:         100,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Class:        "Warrior",
		Race:         "zjeb",
		Subrace:      "zjeb",
		Session:      1,
		Items:        []int{69, 1337, 2137},
	})
	srv.players = append(srv.players, &Player{
		Id:           3,
		Name:         "Test3",
		Level:        131,
		Health:       1,
		Mana:         69,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Class:        "Warrior",
		Race:         "zjeb",
		Subrace:      "zjeb",
		Session:      1,
		Items:        []int{69, 1337, 2137},
	})

	srv.items = append(srv.items, &Item{
		Id:           1,
		Name:         "Test1",
		Description:  "Test description",
		Ability:      "test ability",
		Rarity:       4,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Quantity:     53,
	})
	srv.items = append(srv.items, &Item{
		Id:           2,
		Name:         "Test2",
		Description:  "Test description",
		Ability:      "test ability",
		Rarity:       4,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Quantity:     53,
	})
	srv.items = append(srv.items, &Item{
		Id:           3,
		Name:         "Test3",
		Description:  "Test description",
		Ability:      "test ability",
		Rarity:       4,
		Strength:     10,
		Endurance:    10,
		Perception:   10,
		Intelligence: 10,
		Agility:      10,
		Accuracy:     10,
		Charisma:     10,
		Quantity:     53,
	})
}

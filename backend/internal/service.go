package internal

import (
	"database/sql"

	"github.com/pkg/errors"
)

type ObjectType string

const (
	PlayerType ObjectType = "player"
	ItemType   ObjectType = "item"
	BagType    ObjectType = "bag"
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

func (srv *Service) ResetObjects(types ...ObjectType) error {
	for _, t := range types {
		switch t {
		case PlayerType:
			players, err := srv.GetPlayersFromDB()
			if err != nil {
				return errors.Wrap(err, "ResetObjects: GetPlayersFromDB")
			}
			srv.players = players
		case ItemType:
			items, err := srv.GetItemsFromDB()
			if err != nil {
				return errors.Wrap(err, "ResetObjects: GetItemsFromDB")
			}
			srv.items = items
		case BagType:
			bag, err := srv.GetBagFromDB()
			if err != nil {
				return errors.Wrap(err, "ResetObjects: GetBagFromDB")
			}
			srv.bag = bag
		default:
			return errors.New("unknown object type")
		}
	}
	return nil
}

func (srv *Service) PlayerExists(id int) bool {
	for _, player := range srv.players {
		if player.Id == id {
			return true
		}
	}
	return false
}

func (srv *Service) ItemExists(id int) bool {
	for _, item := range srv.items {
		if item.Id == id {
			return true
		}
	}
	return false
}

func (srv *Service) AlcoholLevelDown() error {
	query := "UPDATE players SET alcohol_level = alcohol_level - 1 WHERE alcohol_level > 0"
	_, err := srv.db.Exec(query)
	if err != nil {
		return err
	}

	return srv.ResetObjects(PlayerType)
}

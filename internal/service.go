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
	return srv, nil
}

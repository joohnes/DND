package internal

import (
	"database/sql"
)

type Service struct {
	db      *sql.DB
	players []*Player
	items   []*Item
}

func NewService(db *sql.DB) (*Service, error) {
	srv := &Service{db: db}
	players, err := srv.GetPlayersFromDB()
	if err != nil {
		return nil, err
	}
	items, err := srv.GetItemsFromDB()
	if err != nil {
		return nil, err
	}

	srv.items = items
	srv.players = players
	return srv, nil
}

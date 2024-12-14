package internal

import (
	"database/sql"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (srv *Service) PlayerExists(id int) bool {
	query := "SELECT 1 FROM players where id = ?"
	_, err := srv.db.Exec(query, id)
	if err != nil {
		return false
	}

	return true
}

func (srv *Service) ItemExists(id int) bool {
	query := "SELECT 1 FROM items where id = ?"
	_, err := srv.db.Exec(query, id)
	if err != nil {
		return false
	}

	return true
}

func (srv *Service) AlcoholLevelDown() error {
	query := "UPDATE players SET alcohol_level = alcohol_level - 1 WHERE alcohol_level > 0"
	_, err := srv.db.Exec(query)
	return err
}

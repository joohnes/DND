package internal

import "errors"

var (
	ErrNoPlayer    = errors.New("player with given id not found")
	ErrNoItem      = errors.New("item with given id not found")
	ErrEmptyID     = errors.New("id cant be empty")
	ErrIDNotNumber = errors.New("id must be numerical")
)

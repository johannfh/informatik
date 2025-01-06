package game

import "errors"

var ErrEntityKilled = errors.New("An Entity just got killed by another entity. How cruel...")

type Entity interface {
	GetID() int
	SetID(id int)

	Tick(*Game) error
}

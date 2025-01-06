package game

import (
	"errors"
	"time"
)

var ErrEntityKilled = errors.New("An Entity just got killed by another entity. How cruel...")

type Entity interface {
	GetID() int
	SetID(id int)

	Tick(g *Game, deltatime time.Duration) error
}

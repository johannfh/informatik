package game

import (
	"errors"
	"time"
)

var ErrEntityKilled = errors.New("An Entity just got killed by another entity. How cruel...")

type Entity interface {
	Tick(g *Game, deltatime time.Duration) error
	GetID() int
	SetID(id int)
}

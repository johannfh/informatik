package game

import (
	"errors"
	"time"
)

var ErrPlantDriedUp = errors.New("A Plant just died because it did not have enough water. :(")

type Plant interface {
	Entity
	GetHeight() Size
	SetHeight(height Size)
}

func NewTree(height Size) *Tree {
	return &Tree{Height: height}
}

type Tree struct {
	ID int

	Height Size
}

func (t *Tree) GetID() int {
	return t.ID
}
func (t *Tree) SetID(id int) {
	t.ID = id
}

func (t *Tree) Tick(g *Game, deltatime time.Duration) error {
	waterCons := t.getWaterCons(deltatime)
	g.Water -= waterCons

	if g.Water < 0 {
		g.Water = 0
		return ErrPlantDriedUp
	}

	return nil
}

func (t *Tree) getWaterCons(deltatime time.Duration) Water {
	// base consume
	consume := t.Height.Meters()

	// timeframe for one full consume (five seconds)
	timeframe := deltatime.Seconds() / 5

	// water consume: consume / timeframe
	return NewWater(consume * timeframe)
}

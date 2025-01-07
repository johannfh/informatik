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
	ID int `json:"id"`

	Height Size `json:"height"`
}

func (t *Tree) GetID() int {
	return t.ID
}
func (t *Tree) SetID(id int) {
	t.ID = id
}

func (t *Tree) Tick(g *Game, deltatime time.Duration) error {
	waterCons := t.getWaterCons(deltatime)
	g.water -= waterCons

	if g.water < 0 {
		g.water = 0
		return ErrPlantDriedUp
	}

	return nil
}

func (t *Tree) getWaterCons(deltatime time.Duration) float64 {
	// base consume
	consume := t.Height.Meters()

	// timeframe for one full consume (five seconds)
	timeframe := deltatime.Seconds() / 5

	// water consume: consume / timeframe
	return consume * timeframe
}

type Flower struct {
	ID int `json:"id"`

	Height Size   `json:"height"`
	Color  string `json:"color"`
}

func (f *Flower) GetID() int {
	return f.ID
}
func (f *Flower) SetID(id int) {
	f.ID = id
}

func (f *Flower) Tick(g *Game, deltatime time.Duration) error {
	waterCons := f.getWaterCons(deltatime)

	g.water -= waterCons
	if g.water < 0 {
		g.water = 0
		return ErrPlantDriedUp
	}

	return nil
}

func (f *Flower) getWaterCons(deltatime time.Duration) float64 {
	// base consume
	consume := f.Height.Meters()

	// timeframe for one full consume (two seconds)
	timeframe := deltatime.Seconds() / 2

	// water consume: consume / timeframe
	return consume * timeframe
}

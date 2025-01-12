package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Sheep struct {
	Type string   `json:"type"`
	ID   EntityID `json:"id"`

	Pos mathutil.Vector2D `json:"position"`
}

func NewSheep() *Sheep {
	return &Sheep{
		Type: SheepType,
	}
}

func (w *Sheep) Tick(g *Game, dt time.Duration) {
	// TODO
}

func (w *Sheep) GetID() EntityID {
	return w.ID
}

func (w *Sheep) SetID(id EntityID) {
	w.ID = id
}

func (w *Sheep) GetPosition() mathutil.Vector2D {
	return w.Pos
}

func (w *Sheep) SetPosition(pos mathutil.Vector2D) {
	w.Pos = pos
}

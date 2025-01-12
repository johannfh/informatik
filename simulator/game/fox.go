package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Fox struct {
	Type string   `json:"type"`
	ID   EntityID `json:"id"`

	Pos mathutil.Vector2D `json:"position"`
}

func NewFox() *Fox {
	return &Fox{
		Type: FoxType,
	}
}

func (w *Fox) Tick(g *Game, dt time.Duration) {
	// TODO
}

func (w *Fox) GetID() EntityID {
	return w.ID
}

func (w *Fox) SetID(id EntityID) {
	w.ID = id
}

func (w *Fox) GetPosition() mathutil.Vector2D {
	return w.Pos
}

func (w *Fox) SetPosition(pos mathutil.Vector2D) {
	w.Pos = pos
}

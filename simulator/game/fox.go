package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Fox struct {
	ID  EntityID
	Pos mathutil.Vector2D
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

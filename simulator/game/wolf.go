package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Wolf struct {
	ID  EntityID
	Pos mathutil.Vector2D
}

func (w *Wolf) Tick(g *Game, dt time.Duration) {
	// TODO
}

func (w *Wolf) GetID() EntityID {
	return w.ID
}

func (w *Wolf) SetID(id EntityID) {
	w.ID = id
}

func (w *Wolf) GetPosition() mathutil.Vector2D {
	return w.Pos
}

func (w *Wolf) SetPosition(pos mathutil.Vector2D) {
	w.Pos = pos
}

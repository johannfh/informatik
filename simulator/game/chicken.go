package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Chicken struct {
	ID  EntityID
	Pos mathutil.Vector2D
}

func (w *Chicken) Tick(g *Game, dt time.Duration) {
	// TODO
}

func (w *Chicken) GetID() EntityID {
	return w.ID
}

func (w *Chicken) SetID(id EntityID) {
	w.ID = id
}

func (w *Chicken) GetPosition() mathutil.Vector2D {
	return w.Pos
}

func (w *Chicken) SetPosition(pos mathutil.Vector2D) {
	w.Pos = pos
}

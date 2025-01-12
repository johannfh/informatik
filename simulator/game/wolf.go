package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type Wolf struct {
	Type string   `json:"type"`
	ID   EntityID `json:"id"`

	Pos mathutil.Vector2D `json:"position"`
}

func NewWolf() *Wolf {
	return &Wolf{
		Type: WolfType,
	}
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

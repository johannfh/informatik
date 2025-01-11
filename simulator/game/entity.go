package game

import (
	"time"

	"github.com/johannfh/informatik/simulator/mathutil"
)

type EntityID int

type Entity interface {
	Tick(g *Game, deltatime time.Duration)
	GetID() EntityID
	SetID(id EntityID)
	GetPosition() mathutil.Vector2D
	SetPosition(pos mathutil.Vector2D)
}

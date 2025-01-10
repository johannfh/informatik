package game

import "time"

type Wolf struct{}

func (w *Wolf) Tick(g *Game, dt time.Duration) {}

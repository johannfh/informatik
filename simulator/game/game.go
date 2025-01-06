package game

import (
	"fmt"
	"log/slog"
	"time"
)

// Creates a new [Game]
func NewGame(l *slog.Logger, t *time.Ticker, w Water) *Game {
	return &Game{
		logger: l,
		Water:  w,
		Ticker: t,
	}
}

// The top level class of the simulation
type Game struct {
	logger *slog.Logger
	Water  Water
	Ticker *time.Ticker
}

func (g *Game) Start() {
	previousTickTime := time.Now()
	for {
		currentTickTime := <-g.Ticker.C
		deltaTime := currentTickTime.Sub(previousTickTime)

		g.logger.Info(
			"tick calculated",
			"deltatime", fmt.Sprintf("%.3fs", deltaTime.Seconds()),
		)

		previousTickTime = currentTickTime
	}
}

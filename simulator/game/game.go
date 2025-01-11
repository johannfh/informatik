package game

import (
	"log/slog"
	"time"
)

type Game struct {
	water    int
	entities []Entity

	nextID EntityID
	stop   bool
}

func NewGame() *Game {
	return &Game{
		entities: []Entity{},
	}
}

func (g *Game) Start(ticker *time.Ticker) {
	previous := time.Now()
	for {
		if g.stop {
			ticker.Stop()
			break
		}

		now := <-ticker.C
		deltatime := now.Sub(previous)
		g.Tick(deltatime)
	}
}

func (g *Game) Stop() {
	g.stop = true
}

func (g *Game) Reset() {
	g.stop = true
	*g = *NewGame()
}

func (g *Game) Tick(dt time.Duration) {
	start := time.Now()
	for _, entity := range g.entities {
		entity.Tick(g, dt)
	}
	end := time.Now()
	slog.Debug(
		"tick calculated",
		"start", start,
		"end", end,
		"duration", end.Sub(start),
	)
}

func (g *Game) AddWater(water int) {
	g.water += water
}

func (g *Game) GetWater() int {
	return g.water
}

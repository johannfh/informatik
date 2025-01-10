package game

import (
	"log/slog"
	"time"
)

type EntityID int

type Entity interface {
	Tick(g *Game, deltatime time.Duration)
	GetID() EntityID
	SetID(id EntityID)
}

type Game struct {
	water    int
	nextID   EntityID
	ticker   *time.Ticker
	entities []Entity
}

func NewGame(t *time.Ticker) *Game {
	return &Game{
		water:    0,
		entities: []Entity{},

		nextID: 0,
		ticker: t,
	}
}

func (g *Game) Reset() {
	ticker := g.ticker
	*g = *NewGame(ticker)
}

func (g *Game) Start() {
	previous := time.Now()
	for {
		now := <-g.ticker.C
		deltatime := now.Sub(previous)
		g.Tick(deltatime)
	}
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

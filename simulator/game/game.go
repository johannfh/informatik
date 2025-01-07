package game

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/johannfh/informatik/simulator/utils"
)

// Creates a new [Game]
func NewGame(l *slog.Logger, t *time.Ticker, w float64) *Game {
	return &Game{
		logger: l,
		ticker: t,
		nextID: 0,

		entities: make([]Entity, 0),

		// water in liters
		water: w,
	}
}

// The top level class of the simulation
type Game struct {
	// operations
	logger *slog.Logger
	ticker *time.Ticker

	// game state
	entities []Entity
	water    float64
	nextID   int
}

func (g *Game) AddEntity(e Entity) {
	e.SetID(g.nextID)
	g.nextID++

	g.entities = append(g.entities, e)
}

func (g *Game) GetUniqueID() int {
	defer func() { g.nextID++ }()
	return g.nextID
}

func (g *Game) RemoveEntity(id int) {
	idx := slices.IndexFunc(g.entities, func(e Entity) bool {
		return e.GetID() == id
	})

	if idx == -1 {
		return
	}

	g.entities = append(g.entities[:idx], g.entities[idx+1:]...)
}

func (g *Game) CountEntities() int {
	return len(g.entities)
}

func (g *Game) GetWater() float64 {
	return g.water
}
func (g *Game) SetWater(water float64) {
	g.water = water
	if g.water < 0 {
		g.water = 0
	}
}

func (g *Game) Start(ctx context.Context) {
	previousTickTime := time.Now()

	for {
		select {
		case currentTickTime := <-g.ticker.C:
			deltatime := currentTickTime.Sub(previousTickTime)

			g.Tick(deltatime)

			previousTickTime = currentTickTime
		case <-ctx.Done():
			break
		}
	}
}

func (g *Game) Tick(deltatime time.Duration) {
	for _, entity := range g.entities {
		if err := entity.Tick(g, deltatime); err != nil {
			g.logger.Info(err.Error())
			g.RemoveEntity(entity.GetID())
		}
	}

	g.logger.Info(
		"tick calculated",
		"deltatime", fmt.Sprintf("%.3fs", deltatime.Seconds()),
		"water", g.water,
		"entities", g.CountEntities(),
	)
}

func (g *Game) GetEntityCounts() utils.CountMap {
	entityCounts := utils.CountMap{}

	for _, entity := range g.entities {
		switch entity.(type) {
		case *Tree:
			entityCounts.Increment("Plants/Trees")
		case *Flower:
			entityCounts.Increment("Plants/Flowers")
		case *Wolf:
			entityCounts.Increment("Animals/Wolves")
		}
	}

	return entityCounts
}

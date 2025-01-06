package game

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/johannfh/go-utils/assert"
)

// Creates a new [Game]
func NewGame(l *slog.Logger, t *time.Ticker, w Water) *Game {
	return &Game{
		entities: make([]Entity, 0),
		logger:   l,
		Water:    w,
		Ticker:   t,
	}
}

// The top level class of the simulation
type Game struct {
	// operations
	logger *slog.Logger
	Ticker *time.Ticker

	// game state
	entities []Entity
	Water    Water
}

func (g *Game) AddEntity(e Entity) {
	g.entities = append(g.entities, e)
}

func (g *Game) Start() {
	previousTickTime := time.Now()

	for {
		currentTickTime := <-g.Ticker.C
		deltatime := currentTickTime.Sub(previousTickTime)

		g.Tick(deltatime)

		previousTickTime = currentTickTime
	}
}

// TODO: Maybe a tree is nice here for
// better visualizing relations between counts
type EntityCounts map[string]int

func (e EntityCounts) Inc(key string) {
	assert.Assert(key != "" && key != "total", "invalid counter key")
	e["total"]++
	e[key]++
}

func (g *Game) Tick(deltatime time.Duration) {
	ec := EntityCounts{}
	for _, entity := range g.entities {
		if err := entity.Tick(g, deltatime); err != nil {
			g.logger.Info(err.Error())
		}

		switch entity.(type) {
		case *Tree:
			ec.Inc("trees")
		}
	}

	g.logger.Info(
		"tick calculated",
		"deltatime", fmt.Sprintf("%.3fs", deltatime.Seconds()),
		"water", g.Water,
		"entities", ec,
	)
}

package game

import (
	"log/slog"
	"sync"
	"time"

	"github.com/johannfh/informatik/simulator/utils"
)

type Game struct {
	water   utils.Observable[float64]
	waterMu sync.Mutex

	entities   utils.Observable[[]Entity]
	entitiesMu sync.Mutex

	nextID   EntityID
	nextIDMu sync.Mutex

	stop   bool
	stopMu sync.Mutex
}

func NewGame() *Game {
	return &Game{
		entities: utils.NewObservable([]Entity{}),
		water:    utils.NewObservable(0.0),
	}
}

func (g *Game) Start(ticker *time.Ticker) {
	previous := time.Now()
	for {
		g.stopMu.Lock()
		if g.stop {
			g.stopMu.Unlock()
			ticker.Stop()
			break
		}
		g.stopMu.Unlock()

		now := <-ticker.C
		deltatime := now.Sub(previous)
		g.Tick(deltatime)
		previous = now
	}
}

func (g *Game) Stop() {
	g.stopMu.Lock()
	defer g.stopMu.Unlock()
	g.stop = true
}

func (g *Game) Reset() {
	g.stopMu.Lock()
	defer g.stopMu.Unlock()

	g.stop = true
	*g = *NewGame()
}

func (g *Game) getUniqueID() EntityID {
	g.nextIDMu.Lock()
	defer g.nextIDMu.Unlock()
	nextID := g.nextID
	g.nextID++
	return nextID
}

func (g *Game) Tick(dt time.Duration) {
	start := time.Now()
	for _, entity := range g.entities.Get() {
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

func (g *Game) AddEntity(entity Entity) {
	entity.SetID(g.getUniqueID())
	slog.Info("added entity")

	g.entitiesMu.Lock()
	defer g.entitiesMu.Unlock()
	g.entities.Set(append(g.entities.Get(), entity))
}

func (g *Game) OnEntitiesChange(fn utils.ListenerFunc[[]Entity]) {
	g.entitiesMu.Lock()
	defer g.entitiesMu.Unlock()
	g.entities.OnChange(fn)
}

func (g *Game) AddWater(water float64) {
	g.waterMu.Lock()
	defer g.waterMu.Unlock()
	val := g.water.Get() + water
	if val < 0 {
		g.water.Set(0)
	} else {
		g.water.Set(val)
	}
}

func (g *Game) GetWater() float64 {
	g.waterMu.Lock()
	defer g.waterMu.Unlock()
	return g.water.Get()
}

func (g *Game) OnWaterChange(fn utils.ListenerFunc[float64]) {
	g.waterMu.Lock()
	defer g.waterMu.Unlock()
	g.water.OnChange(fn)
}

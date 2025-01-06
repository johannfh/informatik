package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/johannfh/informatik/simulator/game"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ticker := time.NewTicker(500 * time.Millisecond)
	initialWater := 10000.0

	g := game.NewGame(logger, ticker, game.NewWater(initialWater))
	g.AddEntity(game.NewTree(game.NewSize(10)))

	g.Start()
}

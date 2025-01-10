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

	ticker := time.NewTicker(time.Second)
	game := game.NewGame(ticker)

	game.Start()
}

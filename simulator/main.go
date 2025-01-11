package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/johannfh/informatik/simulator/api"
	"github.com/johannfh/informatik/simulator/game"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	game := game.NewGame()
	ticker := time.NewTicker(time.Second)
	// Game is running in the background
	go game.Start(ticker)

	srv := api.NewServer()
	srv.ListenAndServe(":8080")
}

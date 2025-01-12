package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/johannfh/informatik/simulator/api"
	"github.com/johannfh/informatik/simulator/game"
	"github.com/johannfh/informatik/simulator/utils"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := utils.ContextWithLogger(context.Background(), logger)

	game := game.NewGame()
	ticker := time.NewTicker(time.Second)
	// Game is running in the background
	go game.Start(ticker)

	hub := api.NewHub(ctx, game)
	go hub.Run()

	srv := api.NewServer(ctx, hub)
	srv.ListenAndServe(":8080")
}

package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/johannfh/informatik/simulator/api"
	"github.com/johannfh/informatik/simulator/game"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	ctx := context.Background()

	ticker := time.NewTicker(500 * time.Millisecond)
	initialWater := 10.0

	g := game.NewGame(logger, ticker, initialWater)
	g.AddEntity(game.NewTree(game.NewSize(10)))
	g.AddEntity(game.NewWolf(20))

	go g.Start(ctx)

	addr := ":8080"
	go log.Fatal(http.ListenAndServe(addr, api.ApiServer{
		Context: ctx,
		Game:    g,
	}.CreateRouter("/")))

	<-ctx.Done()
}

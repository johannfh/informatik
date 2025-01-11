package api

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/johannfh/informatik/simulator/utils"
)

type Server struct {
	wsHub   *Hub
	context context.Context
	logger  *slog.Logger
}

func NewServer(ctx context.Context, hub *Hub) *Server {
	return &Server{
		wsHub:   hub,
		context: ctx,
		logger:  utils.LoggerFromContext(ctx),
	}
}

func (s *Server) ListenAndServe(addr string) {
	r := chi.NewRouter()

	r.Get("/ws", s.handleWs)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		s.logger.Error("server crashed", "err", err)
		os.Exit(1)
	}
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("new connection")
	serveWs(s.context, s.wsHub, w, r)
}

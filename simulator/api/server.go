package api

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	wsHub *Hub
}

func NewServer() *Server {
	return &Server{
		wsHub: NewHub(),
	}
}

func (s *Server) ListenAndServe(addr string) {
	r := chi.NewRouter()

	r.Get("/ws", s.handleWs)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		slog.Error("server crashed", "err", err)
		os.Exit(1)
	}
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	serveWs(s.wsHub, w, r)
}

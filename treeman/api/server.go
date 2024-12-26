package api

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *slog.Logger
}

func (s Server) CreateRouter(frontendFS fs.FS) (chi.Router, error) {
	r := chi.NewRouter()

	r.Handle("/frontend/*", http.StripPrefix("/frontend/", http.FileServerFS(frontendFS)))

	return r, nil
}

func (s *Server) handleGetFrontent(w http.ResponseWriter, r *http.Request) {}

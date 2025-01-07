package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/johannfh/go-utils/assert"
)

type Server struct {
	Context context.Context
}

func (s Server) CreateRouter(prefix string) chi.Router {
	assert.NotEmpty(prefix, "empty prefix")
	r := chi.NewRouter()

	prefixedRouter := chi.NewRouter()
	prefixedRouter.Mount(prefix, r)
	return r
}

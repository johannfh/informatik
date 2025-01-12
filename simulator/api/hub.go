package api

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/johannfh/informatik/simulator/game"
	"github.com/johannfh/informatik/simulator/utils"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register   chan *Client
	unregister chan *Client

	context context.Context
	logger  *slog.Logger

	game *game.Game
}

func NewHub(ctx context.Context, g *game.Game) *Hub {
	l := utils.LoggerFromContext(ctx)
	hub := &Hub{
		clients: make(map[*Client]bool),

		// broadcast a message to all clients
		broadcast: make(chan []byte),

		// register a new client
		register: make(chan *Client),
		// remove a client registration
		unregister: make(chan *Client),

		context: ctx,
		logger:  l,

		// reference to global game state
		game: g,
	}

	return hub
}

func (h *Hub) Run() {
	h.game.OnWaterChange(func(val float64) {
		res, err := json.Marshal(NewServerGameWaterUpdatedMessage(
			val,
		))
		if err != nil {
			h.logger.Error("failed to encode json", "err", err, "data", res)
			return
		}
		h.broadcast <- res
	})

	h.game.OnEntitiesChange(func(val []game.Entity) {
		res, err := json.Marshal(NewServerGameEntitiesUpdatedMessage(
			val,
		))
		if err != nil {
			h.logger.Error("failed to encode json", "err", err)
			return
		}
		h.broadcast <- res
	})
	for {
		select {
		case client := <-h.register:
			h.SendAllStateToClient(client)
			h.clients[client] = true

		case client := <-h.unregister:
			delete(h.clients, client)
		case msg := <-h.broadcast:
			h.logger.Info("broadcasting message")
			for client := range h.clients {
				select {
				// Try sending the message to the connection
				case client.send <- msg:
				// Close the connection if channel is full
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) SendAllStateToClient(c *Client) {
	water := h.game.GetWater()
	waterMsg, err := json.Marshal(NewServerGameWaterUpdatedMessage(
		water,
	))
	if err != nil {
		h.logger.Error("failed to encode json", "err", err, "data", waterMsg)
		return
	}
	c.send <- waterMsg

	entities := h.game.GetEntities()
	entitiesMsg, err := json.Marshal(NewServerGameEntitiesUpdatedMessage(
		entities,
	))
	if err != nil {
		h.logger.Error("failed to encode json", "err", err)
		return
	}
	c.send <- entitiesMsg
}

package api

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johannfh/informatik/simulator/utils"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Allow all origins
	CheckOrigin: func(r *http.Request) bool {
		// TODO: implement proper checking in respective environments
		return true
	},
}

// Client is the middleman between the websocket connection and the hub
type Client struct {
	// The global websocket hub
	hub *Hub

	// The websocket connection
	conn *websocket.Conn

	// Buffered channel for outbound messages
	send chan []byte

	context context.Context

	logger *slog.Logger
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Error("unexpected websocket close error", "err", err)
			}
			return
		}
		c.logger.Info("new message")
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.logger.Info("received message from websocket connection", "message", string(message))
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			// TODO: Figure out a way to batch messages?

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request) {
	logger := utils.LoggerFromContext(ctx)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("failed to upgrade connection to websocket protocol", "err", err)
		return
	}
	client := &Client{
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte, 256),
		context: ctx,
		logger:  logger,
	}
	client.hub.register <- client
	client.logger.Info("client registered")

	// Allow collection of memory referenced by the caller
	// by doing all work in new goroutines.
	go client.writePump()
	go client.readPump()
}

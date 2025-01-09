package api

import (
	"time"

	"github.com/gorilla/websocket"
)

func CloseConn(conn *websocket.Conn, code int, reason string) error {
	msg := websocket.FormatCloseMessage(code, reason)
	if err := conn.WriteControl(websocket.CloseMessage, msg, time.Now().Add(time.Second)); err != nil {
		return err
	}

	return nil
}

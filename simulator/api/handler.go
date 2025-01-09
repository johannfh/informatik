package api

import (
	"bytes"
	"encoding/json"
	"log/slog"

	"github.com/gorilla/websocket"
)

func (srv *WebsocketServer) handleConn(conn *websocket.Conn, stop func()) {

	mt, msg, err := conn.ReadMessage()
	if err != nil {
		slog.Error("read websocket message error", "err", err)
		stop()
		return
	}

	if mt != websocket.TextMessage {
		CloseConn(conn, websocket.CloseUnsupportedData, "unsupported message type")
		stop()
		return
	}

	r := bytes.NewReader(msg)
	var data map[string]any
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		slog.Error("could not parse", "err", err)
		CloseConn(conn, websocket.CloseUnsupportedData, "could not parse data as object")
		stop()
		return
	}

	messageKind, err := GetMsgKind(data)
	if err != nil {
		slog.Error("could not get message kind", "err", err)
		CloseConn(conn, websocket.CloseUnsupportedData, err.Error())
		stop()
		return
	}

	if !IsValidMsgKind(messageKind) {
		CloseConn(conn, websocket.CloseUnsupportedData, "unknown message kind")
		stop()
		return
	}

	if !IsClientMsgKind(messageKind) {
		CloseConn(conn, websocket.CloseUnsupportedData, "not a client message, must start with 'client.*'")
		stop()
		return
	}

	slog.Info(
		"successfully parsed message",
		"messageKind", messageKind,
		"message", data,
	)

	switch messageKind {
	case MsgKindClientUpdateWater:
		var data ClientGetWaterMsg
		if err := json.Unmarshal(msg, &data); err != nil {
			slog.Error("invalid data format for client.updateWater", "err", err)
			InvalidDataFormat(conn)
			stop()
			return
		}
		srv.handleClientUpdateWater(data)
	}

}

// closes a connection because of invalid data for the expected format
func InvalidDataFormat(conn *websocket.Conn) {
	CloseConn(conn, websocket.CloseUnsupportedData, "invalid data format for this message kind")
}

type ClientGetWaterMsg struct {
	MessageKind string `json:"messageKind"`

	Amount float64 `json:"amount"`
}

func (srv *WebsocketServer) handleClientUpdateWater(msg ClientGetWaterMsg) {
	before := srv.Game.GetWater()
	srv.Game.SetWater(srv.Game.GetWater() + msg.Amount)
	after := srv.Game.GetWater()
	change := after - before

	slog.Info(
		"Handle: client.getWater",
		"before", before,
		"after", after,
		"change", change,
	)
}

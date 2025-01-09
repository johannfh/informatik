package api

import (
	"errors"
	"strings"
)

const (
	// like 'a plan just died'
	MsgKindServerEvent = "server.event"
	// like '{"entities": 10, ...}'
	MsgKindServerStats = "server.stats"

	// restarts the game
	MsgKindClientRestartGame = "client.restartGame"
	// changes the water level in the game
	MsgKindClientUpdateWater = "client.updateWater"
)

// Contains every valid websocket message type
// together with their owner/user (client/server)
var MsgKinds = []string{
	// server sent events
	MsgKindServerEvent,
	MsgKindServerStats,

	// client sent events
	MsgKindClientRestartGame,
	MsgKindClientUpdateWater,
}

func IsValidMsgKind(msgKind string) bool {
	for _, v := range MsgKinds {
		if v == msgKind {
			return true
		}
	}
	return false
}

func IsServerMsgKind(msgKind string) bool {
	return strings.HasPrefix(msgKind, "server.")
}

func IsClientMsgKind(msgKind string) bool {
	return strings.HasPrefix(msgKind, "client.")
}

var ErrMsgKindInvalidType = errors.New("property 'messageKind' is not of type string")

func GetMsgKind(msg map[string]any) (string, error) {
	val := msg["messageKind"]
	if kind, ok := val.(string); ok {
		return kind, nil
	}

	return "", ErrMsgKindInvalidType
}

package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/johannfh/informatik/simulator/game"
)

const (
	MsgClientGameSetWater    = "client.game.water.add"
	MsgClientGameEntitiesAdd = "client.game.entities.add"

	MsgServerGameWaterUpdated    = "server.game.water.updated"
	MsgServerGameEntitiesUpdated = "server.game.entities.updated"
	MsgServerGameEvent           = "server.game.event"

	MsgErrorInvalidMessageFormat = "error.invalidMessageFormat"
	MsgErrorUnknownMessageType   = "error.unknownMessageType"
	MsgErrorUnknownEntityError   = "error.unknownEntityError"
)

type Message map[string]any

func (c *Client) handleMessage(msg []byte) {
	if !json.Valid(msg) {
		c.logger.Error("received invalid json from client", "message", string(msg))
		res, err := json.Marshal(NewErrorInvalidMessageFormatMessage("message not decodable as json"))
		if err != nil {
			c.logger.Error("failed to encode json", "err", err)
			return
		}
		c.send <- res

		// Fatal violation of protocol.
		// Terminate connection immediately.
		c.conn.Close()
		return
	}

	var raw Message
	if err := json.Unmarshal(msg, &raw); err != nil {
		c.logger.Error("failed to decode message", "err", err)
		return
	}

	messageType := raw["messageType"]
	switch messageType {
	case MsgClientGameSetWater:
		var data ClientGameSetWaterMessage
		if err := json.Unmarshal(msg, &data); err != nil {
			c.logger.Error(
				"failed to decode message", "err", err,
				"messageType", messageType, "message", string(msg),
			)
			res, err := json.Marshal(NewErrorInvalidMessageFormatMessage(
				fmt.Sprintf("incorrect format for message type '%s'", messageType),
			))
			if err != nil {
				c.logger.Error("failed to encode json", "err", err)
				return
			}
			c.send <- res
			return
		}
		c.handleAddWater(data)
	case MsgClientGameEntitiesAdd:
		var data ClientGameEntitiesAddMessage
		if err := json.Unmarshal(msg, &data); err != nil {
			res, err := json.Marshal(NewErrorInvalidMessageFormatMessage(
				fmt.Sprintf("incorrect format for message type '%s'", messageType),
			))
			if err != nil {
				c.logger.Error("failed to encode json", "err", err)
				return
			}
			c.send <- res
			return
		}
		c.handleAddEntity(data)
	default:
		c.logger.Error("received unkown message type", "messageType", messageType)
		res, err := json.Marshal(NewErrorUnknownMessageTypeMessage(
			messageType,
		))
		if err != nil {
			slog.Error("failed to encode json", "err", err)
			return
		}
		c.send <- res
	}
}

type ClientGameSetWaterMessage struct {
	MessageType string `json:"messageType"`

	Water float64 `json:"water"`
}

type ServerGameWaterUpdatedMessage struct {
	MessageType string `json:"messageType"`

	Water float64 `json:"water"`
}

func NewServerGameWaterUpdatedMessage(water float64) ServerGameWaterUpdatedMessage {
	return ServerGameWaterUpdatedMessage{
		MessageType: MsgServerGameWaterUpdated,
		Water:       water,
	}
}

func (c *Client) handleAddWater(data ClientGameSetWaterMessage) {
	c.hub.game.AddWater(data.Water)
}

type ClientGameEntitiesAddMessage struct {
	MessageType string `json:"messageType"`

	Entity string `json:"entity"`
}

type ServerGameEntitiesUpdatedMessage struct {
	MessageType string `json:"messageType"`

	Entities []game.Entity `json:"entities"`
}

func NewServerGameEntitiesUpdatedMessage(entities []game.Entity) ServerGameEntitiesUpdatedMessage {
	return ServerGameEntitiesUpdatedMessage{
		MessageType: MsgServerGameEntitiesUpdated,
		Entities:    entities,
	}
}

func (c *Client) handleAddEntity(data ClientGameEntitiesAddMessage) {
	switch strings.ToLower(data.Entity) {
	case "wolf":
		c.hub.game.AddEntity(game.NewWolf())
	case "sheep":
		c.hub.game.AddEntity(game.NewSheep())
	case "fox":
		c.hub.game.AddEntity(game.NewFox())
	case "chicken":
		c.hub.game.AddEntity(game.NewChicken())
	default:
		res, err := json.Marshal(NewErrorUnknownEntityMessage(
			data.Entity,
		))
		if err != nil {
			c.logger.Error("failed to encode json", "err", err)
			return
		}
		c.send <- res
	}
}

type ErrorMessage struct {
	MessageType string `json:"messageType"`

	Message string `json:"message"`
}

func NewErrorUnknownMessageTypeMessage(messageType any) ErrorMessage {
	return ErrorMessage{
		MessageType: MsgErrorUnknownMessageType,
		Message: fmt.Sprintf(
			"unknown message type: '%v'", messageType,
		),
	}
}

func NewErrorInvalidMessageFormatMessage(msg string) ErrorMessage {
	return ErrorMessage{
		MessageType: MsgErrorInvalidMessageFormat,
		Message:     msg,
	}
}

func NewErrorUnknownEntityMessage(entity string) ErrorMessage {
	return ErrorMessage{
		MessageType: MsgErrorUnknownEntityError,
		Message: fmt.Sprintf(
			"unknown entity: '%s'", entity,
		),
	}
}

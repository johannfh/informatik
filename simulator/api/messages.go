package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

const (
	MsgClientGameSetWater = "client.game.water.add"

	MsgServerGameWaterUpdated = "server.game.water.updated"
	MsgServerGameEvent        = "server.game.event"

	MsgErrorInvalidMessageFormat = "error.invalidMessageFormat"
	MsgErrorUnknownMessageType   = "error.unknownMessageType"
)

type Message map[string]any

func (c *Client) handleMessage(msg []byte) {
	if !json.Valid(msg) {
		c.logger.Error("received invalid json from client")
		res, err := json.Marshal(NewErrorInvalidMessageFormatMessage("message not decodable as json"))
		if err != nil {
			c.logger.Error("failed to encode json", "err", err, "data", res)
			return
		}
		c.send <- res
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
				c.logger.Error("failed to encode json", "err", err, "data", res)
				return
			}
			c.send <- res
			return
		}
		c.handleAddWater(data)
	default:
		c.logger.Error("received unkown message type", "messageType", messageType)
		res, err := json.Marshal(NewErrorUnknownMessageTypeMessage(
			messageType,
		))
		if err != nil {
			slog.Error("failed to encode json", "err", err, "data", res)
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

type ErrorMessage struct {
	MessageType string `json:"messageType"`

	Message string `json:"message"`
}

func NewErrorUnknownMessageTypeMessage(messageType any) ErrorMessage {
	return ErrorMessage{
		MessageType: MsgErrorUnknownMessageType,
		Message:     fmt.Sprintf("unknown message type: '%v'", messageType),
	}
}

func NewErrorInvalidMessageFormatMessage(msg string) ErrorMessage {
	return ErrorMessage{
		MessageType: MsgErrorInvalidMessageFormat,
		Message:     msg,
	}
}

package handler

import (
	"github.com/szsascha/go-realtime-server/network/connection"
	"github.com/szsascha/go-realtime-server/network/message"
)

type Handler struct {
}

func (handler *Handler) Initialize(connection *connection.Connection) {

}

func (handler *Handler) HandleMessage(connection *connection.Connection, request message.Message) bool {
	connection.Send(request)
	return true
}

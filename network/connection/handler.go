package connection

import "github.com/szsascha/go-realtime-server/network/message"

type Handler interface {
	Initialize(connection *Connection)
	HandleMessage(connection *Connection, request message.Message) bool
}

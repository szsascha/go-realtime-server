package main

import (
	"github.com/szsascha/go-realtime-server/echo/handler"
	"github.com/szsascha/go-realtime-server/echo/message"
	"github.com/szsascha/go-realtime-server/network/server"
)

func main() {
	server.Serve(
		":8080",
		&handler.Handler{},
		message.MessageTypes,
	)
}

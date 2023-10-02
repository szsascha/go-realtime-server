package server

import (
	"fmt"
	"net/http"

	"github.com/szsascha/go-realtime-server/network/connection"
	"github.com/szsascha/go-realtime-server/network/message"
)

func Serve(addr string, messageHandler connection.Handler, messageTypes map[string]message.Message) {
	fmt.Println("Register message types:")
	for k, v := range messageTypes {
		fmt.Println(k)
		message.MessageTypes[k] = v
	}

	fmt.Println()
	fmt.Println("Serve on host " + addr)
	go connection.StartPool()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connection.HandleHttp(w, r, messageHandler)
	})
	http.ListenAndServe(addr, nil)
}

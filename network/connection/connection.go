package connection

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/szsascha/go-realtime-server/network/message"

	"github.com/gorilla/websocket"
)

type Connection struct {
	HostId    string
	websocket *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleHttp(w http.ResponseWriter, r *http.Request, handler Handler) {
	newConnection, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	connectionPool.Register <- newConnection

	go newConnection.handle(handler)
}

func (connection *Connection) handle(handler Handler) {
	defer connection.Close()
	handler.Initialize(connection)
	for body := connection.read(); body != nil; body = connection.read() {
		request, err := message.Deserialize(body)
		if err != nil {
			log.Println(err)
			continue
		}

		if !handler.HandleMessage(connection, request) {
			return
		}
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) (*Connection, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	ip, _, _ := net.SplitHostPort(ws.RemoteAddr().String())
	hasher := sha256.New()
	hasher.Write([]byte(ip))
	newConnection := &Connection{
		HostId:    base64.URLEncoding.EncodeToString(hasher.Sum(nil)),
		websocket: ws,
	}
	if err != nil {
		log.Println(err)
		return newConnection, err
	}
	return newConnection, nil
}

func (connection *Connection) read() []byte {
	_, p, err := connection.websocket.ReadMessage()
	if err != nil {
		log.Println(err)
		return nil
	}
	return p
}

func (connection *Connection) write(content []byte) error {
	w, err := connection.websocket.NextWriter(websocket.TextMessage)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := w.Write(content); err != nil {
		fmt.Println(err)
		return err
	}
	if err := w.Close(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func forEachConnection(consumer func(connection *Connection, status bool)) {
	for connection, status := range connectionPool.Connections {
		consumer(connection, status)
	}
}

func (connection *Connection) Send(msg message.Message) error {
	serialized, err := message.Serialize(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return connection.write(serialized)
}

func (connection *Connection) Broadcast(msg message.Message) {
	forEachConnection(func(connection *Connection, status bool) {
		connection.Send(msg)
	})
}

func (connection *Connection) Close() {
	forEachConnection(func(c *Connection, status bool) {
		if c == connection {
			connectionPool.Unregister <- c
		}
	})

	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}

	fmt.Println("Connection closed")
	connection.websocket.Close()
}

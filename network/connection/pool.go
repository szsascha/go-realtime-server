package connection

import "fmt"

var connectionPool = newPool()

type pool struct {
	Register    chan *Connection
	Unregister  chan *Connection
	Connections map[*Connection]bool
}

func newPool() *pool {
	return &pool{
		Register:    make(chan *Connection),
		Unregister:  make(chan *Connection),
		Connections: make(map[*Connection]bool),
	}
}

func StartPool() {
	for {
		select {
		case connection := <-connectionPool.Register:
			connectionPool.Connections[connection] = true
			fmt.Println("Size of Connection Pool: ", len(connectionPool.Connections))
			break
		case connection := <-connectionPool.Unregister:
			delete(connectionPool.Connections, connection)
			fmt.Println("Size of Connection Pool: ", len(connectionPool.Connections))
			break
		}
	}
}

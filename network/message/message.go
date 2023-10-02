package message

type Message interface {
	GetType() string
	GetBody() []byte
}

// Injected by implementation (see server.go)
var MessageTypes = map[string]Message{}

package message

import "github.com/szsascha/go-realtime-server/network/message"

var MessageTypes = map[string]message.Message{
	"echo": &Echo{},
}

package message

import (
	"encoding/json"
)

func Serialize(message Message) ([]byte, error) {
	return json.Marshal(message)
}

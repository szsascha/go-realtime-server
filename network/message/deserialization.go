package message

import (
	"encoding/json"
	"errors"
	"reflect"
)

func Deserialize(data []byte) (Message, error) {
	messageType, err := extractMessageType(data)
	if err != nil {
		return nil, err
	}

	message, err := findAndCreateMessage(messageType)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func extractMessageType(data []byte) (string, error) {
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return "", err
	}

	if messageType, ok := obj["type"].(string); ok {
		return messageType, nil
	}

	return "", errors.New("Type not found")
}

func findAndCreateMessage(messageType string) (Message, error) {
	origin, found := MessageTypes[messageType]
	if !found {
		return nil, errors.New("Message of type '" + messageType + "' not found!")
	}
	foundMessageType := reflect.TypeOf(origin).Elem()
	newMessage := reflect.New(foundMessageType).Interface().(Message)
	return newMessage, nil
}

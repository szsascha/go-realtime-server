# go realtime server

As part of a larger non-public, private project I needed to create a websocket realtime server in go. The server accepts messages of different types and executes different actions depending on the given message type.

I made the realtime server part of the project modular so I could also use it for other projects. This repository is an extracted result of this modular realtime server.


## Requirements

- Go version >= 1.21.0


## Setup

1. Clone this repo
2. run `go run .` from the main directory of this repo

## Usage

Once the server has started, just send websocket requests to `localhost:8080/ws` with your defined messages.

For example:
```
{
    "type": "echo",
    "body": "test echo message"
}
```

The implemented messages should be defined in the echo/message/message.go map of MessageTypes with the reference to their struct.

## Functionality

- Handling of websocket connection with connection pooling
- Create own implementation of your realtime processes
- Json serialization of messages
- Implemented echo server as an example


## Disclaimer

This project is not intended to be used in production out-of-the-box.
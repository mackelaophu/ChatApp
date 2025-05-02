package websocket

import "github.com/olahol/melody"

type WebSocket struct {
	M melody.Melody
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		M: *melody.New(),
	}
}

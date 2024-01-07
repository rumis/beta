package entity

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientMetadata struct {
	ID string
}

// ClientConnection client connection
type ClientConnection struct {
	mu sync.Mutex

	ClientId string
	Conn     *websocket.Conn
}

// WriteMessage 发送消息
func (c *ClientConnection) WriteMessage(msg []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Conn.WriteMessage(websocket.TextMessage, msg)
}

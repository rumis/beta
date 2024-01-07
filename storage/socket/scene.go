package socket

import (
	"encoding/json"
	"sync"

	"github.com/rumis/beta/entity"
)

// Scenes
type Scenes struct {
	Conns map[string]*entity.ClientConnection

	su sync.Mutex
}

var scenesInst *Scenes
var scenesOnce sync.Once

// NewScenes 新WS服务对象
func NewScenes() *Scenes {
	scenesOnce.Do(func() {
		scenesInst = &Scenes{
			Conns: make(map[string]*entity.ClientConnection),
		}
	})
	return scenesInst
}

// Register 添加新链接对象
func (c *Scenes) Register(conn *entity.ClientConnection) {
	c.su.Lock()
	defer c.su.Unlock()

	c.Conns[conn.ClientId] = conn
}

// ConnUnRegister 删除链接对象
func (c *Scenes) UnRegister(conn *entity.ClientConnection) {
	c.su.Lock()
	defer c.su.Unlock()

	delete(c.Conns, conn.ClientId)
}

// ChatChunkEmit send the chat message
func (c *Scenes) ChatChunkEmit(chunk entity.ChatResponseChunk) error {
	buf, err := json.Marshal(chunk)
	if err != nil {
		return err
	}

	c.su.Lock()
	defer c.su.Unlock()

	for _, conn := range c.Conns {
		err = conn.WriteMessage(buf)
		if err != nil {
			return err
		}
	}
	return nil
}

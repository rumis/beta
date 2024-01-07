package socket

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rumis/beta/entity"
	"github.com/rumis/beta/logger"
	"github.com/rumis/beta/pkg/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	HandshakeTimeout: time.Second * 10,
}

// Upgrade update to web socket protocol
func Upgrade(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(context.Background(), "ws.upgrade", "upgrade error,[error:]%+v", err)
		return
	}
	defer c.Close()

	// 本链接的元数据定义
	connMeta := entity.ClientMetadata{
		ID: utils.UUID(),
	}

	// 注册链接
	NewScenes().Register(&entity.ClientConnection{
		ClientId: connMeta.ID,
		Conn:     c,
	})

	for {
		_, message, err := c.ReadMessage()
		// 链接断开
		if websocket.IsUnexpectedCloseError(err) {
			logger.Debug(context.Background(), "ws.close", "%+v", connMeta)
			// 注销链接
			NewScenes().UnRegister(&entity.ClientConnection{
				ClientId: connMeta.ID,
			})
			return
		}
		if err != nil && !websocket.IsUnexpectedCloseError(err) {
			logger.Error(context.Background(), "ws.readmessage", "readmessage error,[error:]%+v", err)
			continue
		}

		// 记录收到的每条信息
		logger.Info(context.Background(), "ws.message.client", "%s", string(message))
	}
}

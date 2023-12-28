package chat

import (
	"context"

	"github.com/rumis/beta/entity"
)

// ChatService interface for ai chat
type ChatService interface {
	Start(ctx context.Context, promp string) (entity.ChatResponse, error)
}

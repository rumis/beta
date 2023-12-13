package chat

import "context"

// Chat interface for ai chat
type Chat interface {
	Start(ctx context.Context) error
}

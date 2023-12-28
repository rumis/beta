package chat

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rumis/beta/entity"
	"github.com/rumis/beta/enum"
	"github.com/rumis/beta/storage"
)

// ChatCompletionStreamService is the client API for Chat service which return data with stream.
type ChatCompletionStreamService struct {
}

var chatStreamService ChatService
var chatStreamServiceOnce sync.Once

// NewChatCompletionStreamService creates a new ChatCompletionStreamService.
func NewChatCompletionStreamService() ChatService {
	chatStreamServiceOnce.Do(func() {
		chatStreamService = &ChatCompletionStreamService{}
	})
	return chatStreamService
}

// Start is a function to start chat with prompt
func (s *ChatCompletionStreamService) Start(ctx context.Context, prompt string) (entity.ChatResponse, error) {

	req := entity.ChatCompletionRequest{
		Model: enum.GPT3Dot5Turbo,
		Messages: []entity.ChatCompletionMessage{
			// {Role: "system", Content: "You are a helpful assistant"},
			{Role: "user", Content: prompt},
		},
		Stream: true,
	}

	chatId, respCh, err := storage.ChatStreamCompletion(ctx, req)
	if err != nil {
		return entity.ChatResponse{}, err
	}

	go func() {
		for {
			chunk, ok := <-respCh
			if !ok {
				fmt.Println("svc break")
				break
			}
			// fmt.Println("chunk:", chunk)
			for _, choice := range chunk.Choices {
				if choice.FinishReason == "stop" {
					fmt.Println("")
					fmt.Println("stop")
					return
				}
				fmt.Print(choice.Delta.Content)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	return entity.ChatResponse{
		ID: chatId,
	}, nil
}

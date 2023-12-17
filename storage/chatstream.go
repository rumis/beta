package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/rumis/beta/entity"
	"github.com/rumis/ray"
	"github.com/spf13/viper"
)

// ChatStreamCompletion is the client API for Chat service.
func ChatStreamCompletion(ctx context.Context, in entity.ChatCompletionRequest) (entity.ChatCompletionResponse, error) {
	var rsp entity.ChatCompletionResponse

	inBuf, err := json.Marshal(in)
	if err != nil {
		return entity.ChatCompletionResponse{}, err
	}
	opts := ray.NewOptions(
		ray.WithBodyS(string(inBuf)),
		ray.WithURL(viper.GetString("openai.chatHost")+"/v1/chat/completions"),
		ray.WithHeader(map[string]string{
			"Authorization": "Bearer " + viper.GetString("openai.sk"),
			"Content-Type":  "application/json",
			"Aceept":        "text/event-stream",
			"Connection":    "keep-alive",
		}),
	)
	err = ray.DoStream(opts, func(r *bufio.Reader) error {
		line, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Println(string(line))
		return nil
	})
	if err != nil {
		return entity.ChatCompletionResponse{}, err
	}

	return rsp, nil
}

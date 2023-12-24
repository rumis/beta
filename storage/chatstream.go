package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

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
		ray.WithMethod("POST"),
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
		for {
			line, err := r.ReadString('\n')
			if err != nil && err != io.EOF {
				return err
			}
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimPrefix(line, "\n")

			if !strings.HasPrefix(line, "data: ") {
				// 数据不太合规，忽略
				time.Sleep(time.Millisecond * 10)
				continue
			}
			line = strings.TrimPrefix(line, "data: ")
			fmt.Println(line)
			if line == "[DONE]" {
				fmt.Println("stream done")
				return nil // done
			}
		}
		return nil
	})
	if err != nil {
		return entity.ChatCompletionResponse{}, err
	}

	return rsp, nil
}

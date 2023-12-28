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
func ChatStreamCompletion(ctx context.Context, in entity.ChatCompletionRequest) (string, chan entity.ChatCompletionChunkResponse, error) {
	rspCh := make(chan entity.ChatCompletionChunkResponse)

	inBuf, err := json.Marshal(in)
	if err != nil {
		return "", rspCh, err
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
	// wait for the first chat response
	firstCh := make(chan entity.ChatResponseChunkMetadata)

	go func() {
		err = ray.DoStream(opts, func(r *bufio.Reader) error {
			firstResp := true
			for {
				line, err := r.ReadString('\n')
				if err != nil && err != io.EOF {
					return err
				}
				line = strings.TrimSuffix(line, "\n")
				line = strings.TrimPrefix(line, "\n")
				if !strings.HasPrefix(line, "data: ") {
					// response invalid
					time.Sleep(time.Millisecond * 10)
					continue
				}
				line = strings.TrimPrefix(line, "data: ")
				// fmt.Println("data:", line)
				if line == "[DONE]" {
					close(rspCh)
					close(firstCh)
					return nil // done
				}
				// parse chunk data
				var chunk entity.ChatCompletionChunkResponse
				err1 := json.Unmarshal([]byte(line), &chunk)
				if err1 != nil {
					// response invalid
					time.Sleep(time.Millisecond * 10)
					continue
				}
				if firstResp {
					firstResp = false
					firstCh <- entity.ChatResponseChunkMetadata{
						ID:        chunk.ID,
						ErrString: "",
					}
				}
				rspCh <- chunk
			}
		})
		if err != nil {
			firstCh <- entity.ChatResponseChunkMetadata{
				ID:        "",
				ErrString: err.Error(),
			}
		}
	}()

	chunkMetadata := <-firstCh
	if chunkMetadata.ErrString != "" {
		return "", rspCh, fmt.Errorf(chunkMetadata.ErrString)
	}
	return chunkMetadata.ID, rspCh, nil
}

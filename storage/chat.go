package storage

import (
	"context"

	"github.com/rumis/beta/entity"
	"github.com/rumis/ray"
	"github.com/spf13/viper"
)

// ChatCompletion is the client API for Chat service.
func ChatCompletion(ctx context.Context, in entity.ChatCompletionRequest) (entity.ChatCompletionResponse, error) {
	var rsp entity.ChatCompletionResponse
	err := ray.PostRawJson(ctx, viper.GetString("openai.chatHost")+"/v1/chat/completions", in, &rsp, nil, map[string]string{
		"Authorization": "Bearer " + viper.GetString("openai.sk"),
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rumis/beta/entity"
	"github.com/rumis/beta/enum"

	testBoot "github.com/rumis/beta/boot/test"
)

func TestCahtCompletion(t *testing.T) {

	testBoot.BootInit("/home/workspace/beta/config_local/config.toml")

	req := entity.ChatCompletionRequest{
		Model: enum.GPT3Dot5Turbo,
		Messages: []entity.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant"},
			{Role: "user", Content: "Hello"},
		},
	}

	rsp, err := ChatCompletion(context.Background(), req)

	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(rsp, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}

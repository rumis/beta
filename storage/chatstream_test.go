package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/rumis/beta/entity"
	"github.com/rumis/beta/enum"

	testBoot "github.com/rumis/beta/boot/test"
)

func TestCahtStreamCompletion(t *testing.T) {

	testBoot.BootInit("/home/workspace/beta/config_local/config.toml")

	req := entity.ChatCompletionRequest{
		Model: enum.GPT3Dot5Turbo,
		Messages: []entity.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant"},
			{Role: "user", Content: "hello"},
		},
		Stream: true,
	}

	rsp, err := ChatStreamCompletion(context.Background(), req)

	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(rsp, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}

func TestXxx(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println("hello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long texthello - long text:", i)
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println("done")
}

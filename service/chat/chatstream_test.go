package chat

import (
	"context"
	"fmt"
	"testing"
	"time"

	testBoot "github.com/rumis/beta/boot/test"
)

func TestChatStream(t *testing.T) {

	testBoot.BootInit("/home/workspace/beta/config_local/config.toml")

	chatSvc := NewChatCompletionStreamService()

	// prompt := "hello"
	prompt := "how to build a microservice in golang"
	chatRsp, err := chatSvc.Start(context.Background(), prompt)

	if err != nil {
		panic(err)
	}

	fmt.Println("chat:", chatRsp.ID)

	time.Sleep(time.Second * 10)

}

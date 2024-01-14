package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	chatBoot "github.com/rumis/beta/boot/chat"
	"github.com/rumis/beta/entity"
	chatService "github.com/rumis/beta/service/chat"
	"github.com/rumis/beta/storage/socket"
)

func main() {

	// initialize boot
	chatBoot.BootInit()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	e.Any("/upgrade", func(c echo.Context) error {
		socket.Upgrade(c.Response(), c.Request())
		return nil
	})

	e.POST("/chat", func(c echo.Context) error {
		req := entity.ChatRequest{}
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		resp, err := chatService.NewChatCompletionStreamService().Start(context.Background(), req.Prompt, req.ChatID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})

	// 文件服务
	e.Static("/", viper.GetString("public.root"))

	// Start server
	e.Logger.Fatal(e.Start(":7323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, viper.GetString("public.root"))
}

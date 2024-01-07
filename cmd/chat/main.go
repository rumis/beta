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

	// e.GET("/", func(c echo.Context) error {
	// 	// 解析指定文件生成模板对象
	// 	tmpl := template.New("chat.liumurong.org")
	// 	tmpl, err := tmpl.Parse(tpl.TemplateChat)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// 利用给定数据渲染模板，并将结果写入w
	// 	err = tmpl.Execute(c.Response(), "chat.liumurong.org")
	// 	return err
	// })

	// 文件服务
	e.Static("/", "/home/workspace/beta/public")

	// Start server
	e.Logger.Fatal(e.Start(":7323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, viper.GetString("openai.sk"))
}

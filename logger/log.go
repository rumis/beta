package logger

import (
	"context"
	"fmt"
)

type Logger func(ctx context.Context, tag string, args string, v ...interface{})

// Info 信息
func Info(ctx context.Context, tag string, args string, v ...interface{}) {
	fmt.Println("info,", tag, ",", fmt.Sprintf(args, v...))
}

// Warn 警告
func Warn(ctx context.Context, tag string, args string, v ...interface{}) {
	fmt.Println("warn,", tag, ",", fmt.Sprintf(args, v...))
}

// Error 错误
func Error(ctx context.Context, tag string, args string, v ...interface{}) {
	fmt.Println("error,", tag, ",", fmt.Sprintf(args, v...))
}

// Debug 调试
func Debug(ctx context.Context, tag string, args string, v ...interface{}) {
	fmt.Println("debug,", tag, ",", fmt.Sprintf(args, v...))
}

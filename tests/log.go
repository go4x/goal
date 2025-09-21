package tests

import "github.com/go4x/logx"

func NewLog() logx.Logger {
	logx.Init(&logx.LoggerConfig{
		Level:        "debug",
		Format:       "text",
		Dir:          "logs",
		MaxAge:       0,
		LogInConsole: true,
	})
	return logx.GetLogger()
}

package tests

import "github.com/gophero/logx"

func NewLog() *logx.Logger {
	return logx.NewLog(&logx.Zap{
		Level:         "debug",
		Prefix:        "",
		Format:        "text",
		Director:      "logs",
		EncodeLevel:   "cap",
		StacktraceKey: "stacktrace",
		MaxAge:        0,
		ShowLine:      true,
		LogInConsole:  true,
	})
}

package logging

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

func InitLogger() {
	rollingIO := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    64, // megabytes
		MaxBackups: 30, // maxBackUps
		MaxAge:     30, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}

	combineIO := io.MultiWriter(rollingIO, os.Stdout)

	// TODO 根据配置文件选择 Text or JSON
	slog.SetDefault(slog.New(slog.NewTextHandler(combineIO, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	})))
}

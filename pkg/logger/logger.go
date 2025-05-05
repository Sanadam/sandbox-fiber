package logger

import (
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"log/slog"
	"os"
	"sanadam/sandbox-fiber/config"
)

func NewLogger() *slog.Logger {
	zerologLogger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})
	logger := slog.New(slogzerolog.Option{
		Level:  config.NewLogConfig().Level,
		Logger: &zerologLogger,
	}.NewZerologHandler())

	return logger
}

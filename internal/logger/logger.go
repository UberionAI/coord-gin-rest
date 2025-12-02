package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(logLevel string) {
	zerolog.TimeFieldFormat = time.RFC3339

	// Устанавливаем уровень логирования
	level := parseLogLevel(logLevel)
	zerolog.SetGlobalLevel(level)

	// Красивый вывод для консоли
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006/01/02 15:04:05",
	})
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warning", "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.WarnLevel
	}
}

func Get() *zerolog.Logger {
	return &log.Logger
}

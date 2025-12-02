package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Logger wraps zerolog.Logger for application-wide logging
type Logger struct {
	logger zerolog.Logger
}

// New creates a new logger with specified log level
func New(levelStr string) *Logger {
	// Configure human-readable console output
	zerolog.TimeFieldFormat = time.RFC3339
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006/01/02 15:04:05",
	}

	// Parse log level
	level := parseLogLevel(levelStr)
	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Logger{logger: logger}
}

// parseLogLevel converts string to zerolog.Level
func parseLogLevel(levelStr string) zerolog.Level {
	switch strings.ToLower(levelStr) {
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

// Info logs informational message
func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Warn logs warning message
func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

// Error logs error message
func (l *Logger) Error(msg string, err error) {
	l.logger.Error().Err(err).Msg(msg)
}

// Debug logs debug message
func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// WithField adds a field to log entry
func (l *Logger) WithField(key string, value interface{}) *zerolog.Event {
	return l.logger.Info().Interface(key, value)
}

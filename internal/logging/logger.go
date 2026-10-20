package logging

import (
	"context"
	"log/slog"
	"os"
)

// Logger is the centralized interface for logging.
type Logger interface {
	Debug(msg string, attrs ...slog.Attr)
	Info(msg string, attrs ...slog.Attr)
	Warn(msg string, attrs ...slog.Attr)
	Error(msg string, err error, attrs ...slog.Attr)
}

// SlogLogger is an implementation of Logger using slog.
type SlogLogger struct {
	logger *slog.Logger
}

// New creates a new logger with the specified level.
func New(level string) *SlogLogger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stderr, opts)
	logger := slog.New(handler)

	return &SlogLogger{logger: logger}
}

// NewJSON creates a logger with JSON format.
func NewJSON(level string) *SlogLogger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(os.Stderr, opts)
	logger := slog.New(handler)

	return &SlogLogger{logger: logger}
}

// Debug logs a DEBUG-level message.
func (l *SlogLogger) Debug(msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(context.Background(), slog.LevelDebug, msg, attrs...)
}

// Info logs an INFO-level message.
func (l *SlogLogger) Info(msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}

// Warn logs a WARN-level message.
func (l *SlogLogger) Warn(msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(context.Background(), slog.LevelWarn, msg, attrs...)
}

// Error logs an ERROR-level message with an attached error.
func (l *SlogLogger) Error(msg string, err error, attrs ...slog.Attr) {
	if err != nil {
		attrs = append(attrs, slog.Any("error", err))
	}
	l.logger.LogAttrs(context.Background(), slog.LevelError, msg, attrs...)
}

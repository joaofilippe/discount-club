package logger

import (
	"log/slog"
	"os"
	"runtime"
)

type Logger struct {
	Logger *slog.Logger
}

// NewLogger returns a new logger
func NewLogger() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Logger{
		Logger: logger,
	}
}

// Errorf logs an error message with the line number where it occurred
func (l *Logger) Errorf(err error) {
	_, _, line, _ := runtime.Caller(1)

	l.Logger.Error(err.Error(), "line=", line)
}

// Fatalf logs an error message with the line number where it occurred and then panics
func (l *Logger) Fatalf(err error) {
	_, _, line, _ := runtime.Caller(1)

	l.Logger.Error(err.Error(), "line=", line)
	panic(err)
}

// Infof logs an informational message with the line number and file where it occurred
func (l *Logger) Infof(message string) {
	_, file, line, _ := runtime.Caller(1)

	l.Logger.Info(message, "line", line, "file", file)
}
